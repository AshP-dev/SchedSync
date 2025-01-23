//import logo from "./logo.svg";
import "./App.css";
import CardList from "./CardList";
import CreateCard from "./CreateCard";

function App() {
  return (
    <div className='App'>
      <header className='App-header'>
        <main>
          <button>
            <CreateCard />
          </button>
          <CardList />
        </main>
      </header>
    </div>
  );
}

export default App;
