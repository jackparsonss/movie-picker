import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import axios from "axios";

const initialState = {
  toWatch: [],
  watched: [],
};

const getMovies = createAsyncThunk("get/movies", async () => {
  try {
    const req = await axios.get("http://localhost:8080/api/list");
    const req2 = await axios.get("http://localhost:8080/api/list?type=watched");

    return {
      toWatch: req.data,
      watched: req2.data,
    };
  } catch (e) {
    throw new Error(e);
  }
});

const watchMovie = createAsyncThunk("put/watch", async (key) => {
  try {
    const { data } = await axios.put(`http://localhost:8080/api/watch/${key}`);

    return data;
  } catch (e) {
    throw new Error(e);
  }
});

const addMovie = createAsyncThunk("add/movie", async (title) => {
  try {
    const { data } = await axios.post("http://localhost:8080/api/add", {
      title,
    });

    return data;
  } catch (e) {
    throw new Error(e);
  }
});

const moveMovie = createAsyncThunk(
  "move/movie",
  async ({ title, key, bucket }) => {
    try {
      const { data } = await axios.put(
        `http://localhost:8080/api/move/${key}?bucket=${bucket}`,
        {
          title,
        }
      );

      return data;
    } catch (e) {
      throw new Error(e);
    }
  }
);

const deleteMovie = createAsyncThunk(
  "delete/moviie",
  async ({ key, value }, thunkAPI) => {
    try {
      let bucket;
      const watched = thunkAPI.getState().movies.watched;
      let idx = watched.findIndex((m) => m.key === key && m.value === value);
      if (idx !== -1) {
        bucket = "watched";
      }

      const toWatch = thunkAPI.getState().movies.toWatch;
      idx = toWatch.findIndex((m) => m.key === key && m.value === value);
      if (idx !== -1) {
        bucket = "toWatch";
      }

      const { data } = await axios.delete(
        `http://localhost:8080/api/delete/${key}?bucket=${bucket}`
      );

      return data;
    } catch (e) {
      throw new Error(e);
    }
  }
);

const selectMovies = (state) => state.movies;

export const listSlice = createSlice({
  name: "movies",
  initialState,
  reducers: {
    removeFromMovieList: (state, { payload }) => {
      state[payload.type] = state[payload.type].filter(
        (movie) => movie.key !== payload.key
      );
    },
    addToMovieList: (state, { payload }) => {
      state[payload.type].push(payload.data);
    },
  },
  extraReducers: {
    [getMovies.fulfilled]: (state, { payload }) => {
      state.toWatch = payload.toWatch;
      state.watched = payload.watched;
    },
    [watchMovie.fulfilled]: (state, { payload }) => {
      state.toWatch = state.toWatch.filter(
        (movie) => movie.key !== payload.key
      );
      state.watched.push(payload);
    },
    [addMovie.fulfilled]: (state, { payload }) => {
      state.toWatch.push(payload);
    },
    [deleteMovie.fulfilled]: (state, { payload }) => {
      state[payload.bucket] = state[payload.bucket].filter(
        (movie) => movie.key !== payload.key
      );
    },
  },
});

export const { removeFromMovieList, addToMovieList } = listSlice.actions;

export {
  selectMovies,
  getMovies,
  watchMovie,
  addMovie,
  deleteMovie,
  moveMovie,
};

export default listSlice.reducer;
