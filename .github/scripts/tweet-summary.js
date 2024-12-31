const { TwitterApi } = require("twitter-api-v2");
const { execSync } = require("child_process");
require("dotenv").config();

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

const tweetSummary = async () => {
  const summary = getWeeklyCommitSummary();
  try {
    await client.v2.tweet(summary);
    console.log("Tweeted successfully!");
  } catch (error) {
    console.error("Error tweeting:", error);
  }
};

tweetSummary();
