import React from "react";
import ReactDOM from "react-dom";
import "./App.css";
import { InfraMap } from "./components/infra-map";
import { Edge } from "./models/infra-models";

const edges: Edge[] = [];

export const App = () => {
  return (
    <div className="App">
      <header className="App-header">
        <InfraMap edges={edges} />
      </header>
    </div>
  );
};

export default App;
