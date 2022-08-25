import { configureStore } from '@reduxjs/toolkit'
import listReducer from "./slices/list"

export const store = configureStore({
  reducer: {
    movies: listReducer
  },
  devTools: process.env.NODE_ENV !== 'production',
})


