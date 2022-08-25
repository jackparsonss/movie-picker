import './App.css';
import {BrowserRouter, Route, Routes} from "react-router-dom";
import Home from "./screens/Home";
import Picked from "./screens/Picked";

function App() {
  return (
      <BrowserRouter>
          <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/pick" element={<Picked />} />
          </Routes>
      </BrowserRouter>
  );
}

export default App;
