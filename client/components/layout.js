import Sidebar from "./sideBar";
export default function Layout({ children }) {
  return (
    <div className='container'>
      <div className='row'>
        <h1 className='mt-3'>Go watch movie</h1>
        <hr className='mb-3' />
      </div>
      <div className='row'>
        <div className='col-md-2'>
          <Sidebar />
        </div>
        <div className='col-md-10'>{children}</div>
      </div>
    </div>
  );
}
