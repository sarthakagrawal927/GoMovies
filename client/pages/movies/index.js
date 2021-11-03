import Layout from "../../components/layout";
import { useState, useEffect } from "react";
import Link from "next/link";

const moviesData = [
  {
    id: 1,
    title: "The redneck Redemption",
    overview: "Two imprisoned ",
  },
  {
    id: 2,
    title: "The redneck Redemption",
    overview: "Two imprisoned ",
  },
  {
    id: 3,
    title: "The redneck Redemption",
    overview: "Two imprisoned ",
  },
];

export default function Movies() {
  const [movies, setMovies] = useState(moviesData);
  return (
    <Layout>
      <h2> Movies</h2>
      <ul>
        {movies.map((movie, index) => {
          return (
            <li key={index}>
              <Link href={`movies/${movie.id}`}>
                <h4> {movie.title}</h4>
              </Link>

              <p> {movie.overview}</p>
            </li>
          );
        })}
      </ul>
    </Layout>
  );
}
