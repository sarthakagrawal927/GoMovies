import { useRouter } from "next/router";
import Layout from "../../components/layout";

const Movie = () => {
  const router = useRouter();
  const { id } = router.query;
  return (
    <Layout>
      {id}
      <table className='table table-compact table-striped'>
        <thead></thead>
        <tbody>
          <tr>
            <td>Title</td>
            <td>TitleMovie</td>
          </tr>
          <tr>
            <td>Runtime</td>
            <td>RuntimeMovie</td>
          </tr>
        </tbody>
      </table>
    </Layout>
  );
};

export default Movie;
