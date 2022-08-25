import List from "../components/List";
import Pick from "../components/Pick";
import "./Home.css"
import {useEffect} from "react";
import {useDispatch} from "react-redux";
import {getMovies} from "../slices/list";

const Home = () => {
    const dispatch = useDispatch()

    useEffect(() => {
        dispatch(getMovies())
    }, [dispatch])

    return (
        <div className="home">
            <List type="toWatch"/>
            <List type="watched"/>
            <Pick/>
        </div>
    )
}

export default Home