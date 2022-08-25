import { useState } from "react";
import { useDispatch } from "react-redux";
import { addMovie } from "../slices/list";
import "./Add.css";

const Add = () => {
  const [title, setTitle] = useState("");

  const dispatch = useDispatch();

  const handleSubmit = (e) => {
    e.preventDefault();

    dispatch(addMovie(title));
    setTitle("");
  };

  return (
    <form className="add" onSubmit={handleSubmit}>
      <input
        placeholder="movie title"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
      />
      <button type="submit">add</button>
    </form>
  );
};

export default Add;
