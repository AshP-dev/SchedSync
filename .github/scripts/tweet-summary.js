const { TwitterApi } = require("twitter-api-v2");
const { execSync } = require("child_process");

const client = new TwitterApi({
  appKey: process.env.TWITTER_API_KEY,
  appSecret: process.env.TWITTER_API_SECRET,
  accessToken: process.env.TWITTER_ACCESS_TOKEN,
  accessSecret: process.env.TWITTER_ACCESS_SECRET,
});

const getWeeklyCommitSummary = () => {
  const oneWeekAgo = new Date();
  oneWeekAgo.setDate(oneWeekAgo.getDate() - 7);
  const dateString = oneWeekAgo.toISOString().split("T")[0];

  const commitCommand = `git log --since="${dateString}" --oneline | wc -l`;
  const commitCount = execSync(commitCommand).toString().trim();

  const tweetCommand = `git log --since="${dateString}" --pretty=format:"%s"`;
  const tweetMessages = execSync(tweetCommand).toString().trim().split("\n");

  return (
    `Weekly commit summary for ${process.env.GITHUB_REPOSITORY}:\n` +
    `Total commits in the past week: ${commitCount}\n` +
    `Commit messages:\n` +
    tweetMessages.map((msg, index) => `${index + 1}. ${msg}`).join("\n") +
    `\nCheck out our progress: https://github.com/${process.env.GITHUB_REPOSITORY}`
  );
};

const splitMessage = (message) => {
  const maxLength = 250;
  const messages = [];
  while (message.length > maxLength) {
    let splitIndex = message.lastIndexOf(" ", maxLength);
    if (splitIndex === -1) splitIndex = maxLength;
    messages.push(message.slice(0, splitIndex));
    message = message.slice(splitIndex).trim();
  }
  messages.push(message);
  return messages;
};

const tweetSummary = async () => {
  const summary = getWeeklyCommitSummary();
  const messages = splitMessage(summary);
  let replyToId = null;

  for (const message of messages) {
    try {
      const tweet = await client.v2.tweet(
        message,
        replyToId ? { in_reply_to_status_id: replyToId } : {}
      );
      replyToId = tweet.data.id;
      console.log("Tweet sent successfully!");
    } catch (error) {
      console.error("Error sending tweet:", error);
      break;
    }
  }
};

tweetSummary();
