import "./List.css"
import {useSelector} from "react-redux";
import {selectMovies} from "../slices/list";

const List = ({type}) => {
    const movies = useSelector(selectMovies)

    return (
        <div className="list">
            <h3>{type === "watched" ? "Watched" : "To Watch"}</h3>
            <ul>
                {movies[type]?.map(movie => (
                    <li key={movie.key}>{movie.value}</li>
                ))}
            </ul>
        </div>
    )
}

export default List
