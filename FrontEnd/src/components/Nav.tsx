import {Link} from "react-router-dom";


const Nav = (props: { ID: string, setID: (name: string) => void }) => {
    const logout = async () => {
        await fetch('http://localhost:8000/api/logout', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            credentials: 'include',
        });        
    }

    let menu;

        menu = (
            <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item active">
                    <Link to="/login" className="nav-link" onClick={logout}>Login</Link>
                </li>
                <li className="nav-item active">
                    <Link to="/register" className="nav-link" onClick={logout}>Register</Link>
                </li>
                <li className="nav-item active">
                    <Link to="/" className="nav-link" onClick={logout}>Logout</Link>
                </li>
            </ul>
        )
    

    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
                <Link to="/" className="navbar-brand">Home</Link>

                <div>
                    {menu}
                </div>
            </div>
        </nav>
    );
};

export default Nav;
