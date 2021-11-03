import Layout from "../../components/layout";
import Link from "next/link";
export default function Categories() {
  return (
    <Layout>
      <h3> Categories</h3>
      <ul>
        <li>
          <Link href='/by-category/comedy'>Comedy</Link>
        </li>
        <li>
          <Link href='/by-category/drama'>Drama</Link>
        </li>
        <li>
          <Link href='/by-category/thriller'>Thriller</Link>
        </li>
      </ul>
    </Layout>
  );
}
