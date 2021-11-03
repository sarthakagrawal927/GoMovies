import { useRouter } from "next/router";
import Layout from "../../components/layout";

const Movie = () => {
  const router = useRouter();
  const { category } = router.query;
  return (
    <Layout>
      <h3>Category: {category.toUpperCase()}</h3>
    </Layout>
  );
};

export default Movie;
