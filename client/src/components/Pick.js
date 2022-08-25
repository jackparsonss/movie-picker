import "./Pick.css";
import { useNavigate } from "react-router-dom";

const Pick = () => {
  const navigate = useNavigate();
  return (
    <button className="button" onClick={() => navigate("/pick")}>
      Pick a Movie :)
    </button>
  );
};

export default Pick;
