import "./List.css";
import { useDispatch, useSelector } from "react-redux";
import {
  addToMovieList,
  deleteMovie,
  moveMovie,
  removeFromMovieList,
  selectMovies,
} from "../slices/list";
import Add from "./Add";

const List = ({ type }) => {
  const movies = useSelector(selectMovies);

  const dispatch = useDispatch();

  const handleDelete = (key, value) => {
    dispatch(deleteMovie({ key, value }));
  };

  const handleMove = (key, value, type) => {
    dispatch(moveMovie({ title: value, key, bucket: type }));
    dispatch(removeFromMovieList({ type, key }));
    dispatch(
      addToMovieList({
        type: type === "watched" ? "toWatch" : "watched",
        data: { value, key },
      })
    );
  };

  return (
    <div className="list">
      <h3>{type === "watched" ? "Watched" : "To Watch"}</h3>
      <div>
        {movies[type]?.map((movie) => (
          <div className="list__item" key={movie.key}>
            <p>{movie.value}</p>
            <div className="list__items">
              <span
                className={`list__itemWatch ${
                  type === "watched" ? "list__watched" : "list__toWatch"
                }`}
                onClick={() => handleMove(movie.key, movie.value, type)}
              >
                {type === "watched" ? "unwatch" : "watch"}
              </span>
              <span
                className="list__itemDelete"
                onClick={() => handleDelete(movie.key, movie.value)}
              >
                x
              </span>
            </div>
          </div>
        ))}
        {type === "toWatch" && <Add />}
      </div>
    </div>
  );
};

export default List;
