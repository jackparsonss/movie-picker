import "./Picked.scss"
import {useDispatch, useSelector} from "react-redux";
import {getMovies, selectMovies, watchMovie} from "../slices/list";
import React, {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";

const Picked = () => {
    const {toWatch} = useSelector(selectMovies)
    const [movies, setMovies] = useState([]);
    const [currentChoice, setCurrentChoice] = useState('');
    const dispatch = useDispatch()
    const navigate = useNavigate()

    const handleYes = () => {
        const idx = toWatch.findIndex(mov => mov.value === currentChoice)
        dispatch(watchMovie(idx))

        navigate("/")
    }

    useEffect(() => {
        if (toWatch.length === 0) {
            dispatch(getMovies())
        }
        const temp = []
        for (let movie of toWatch) {
            temp.push(movie.value)
        }
        setMovies(temp)
    }, [toWatch.length, movies.length, dispatch])

    return (
        <div className="picked">
            <div className="picked__main">
                {movies.length > 0 &&
                    <RandomPicker items={movies} setCurrentChoice={setCurrentChoice} currentChoice={currentChoice}/>}
                <div className="picked__body">
                    <div className="picked__btns">
                        <p>Would you like to watch this movie?</p>
                        <button className="picked__yes btn" onClick={handleYes}>Yes</button>
                        <button className="picked__no btn" onClick={() => setMovies([])}>No</button>
                    </div>
                </div>
            </div>
        </div>
    )
}
const RandomPicker = ({items, currentChoice, setCurrentChoice}) => {
    let interval = null;
    const intervalDuration = 25;
    const duration = 1500;


    useEffect(() => {
        start()
    }, [])

    const setChoice = () => {
        setCurrentChoice(pickChoice())
    }

    const start = () => {
        clearInterval(interval);
        interval = setInterval(setChoice, intervalDuration);

        setTimeout(() => {
                stop()
        }, duration);
    }

    const stop = () => {
        clearInterval(interval);
    }

    const pickChoice = () => {
        return items[Math.floor(Math.random() * items.length)];
    }

    return (
        <div className="RandomPicker">
            <RandomPickerChoice choice={currentChoice}/>
        </div>
    )
}


class RandomPickerChoice extends React.PureComponent {
    render() {
        const {choice} = this.props;
        const content = choice.trim().length > 0 ? choice : '?';

        return (
            <div className="RandomPicker__choice">
                <span className="RandomPicker__choiceItem">{content}</span>
            </div>
        );
    }
}


export default Picked