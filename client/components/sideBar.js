const SideBar = () => {
  return (
    <nav>
      <ul className='list-group'>
        <li className='list-group-item'>
          <a href='/'>Home</a>
        </li>
        <li className='list-group-item'>
          <a href='/movies'>Movies</a>
        </li>
        <li className='list-group-item'>
          <a href='/by-category'>Genre</a>
        </li>
        <li className='list-group-item'>
          <a href='/admin'>Manage Catalogue</a>
        </li>
      </ul>
    </nav>
  );
};

export default SideBar;
