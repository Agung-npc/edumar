import React, { useState, useContext } from "react";
import { Link, useNavigate } from "react-router-dom";
import { NavDropdown } from "react-bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import AuthContext from "../../store/AuthContext";
import "./style.css";

const Navbar = () => {
  const user = useContext(AuthContext);
  const [sidebar, setSidebar] = useState(false);
  const navigate = useNavigate();

  window.addEventListener("scroll", function () {
    const header = document.querySelector(".header");
    header.classList.toggle("active", window.scrollY > 180);
  });

  const handleLogout = () => {
    navigate("/login");
    localStorage.removeItem("token");
    localStorage.removeItem("answer");
    localStorage.removeItem("result");
  };

  return (
    <>
      <header className="header">
        <div className="container flex">
          <div className="logo">
            <img src="assets/logo.png" alt="" />
          </div>
          <div>
            <ul>
              <li>
                <Link to="/" style={{ textDecoration: "none", color: "black" }}>
                  Home
                </Link>
              </li>
              <li>
                <Link
                  to="/scores-board"
                  style={{ textDecoration: "none", color: "black" }}
                >
                  Scoreboard
                </Link>
              </li>
              {user.token ? (
                <li className="account">
                  <NavDropdown id="navbarScrollingDropdown">
                    <NavDropdown.Item onClick={() => handleLogout()}>
                      Log Out &nbsp;
                    </NavDropdown.Item>
                  </NavDropdown>
                </li>
              ) : null}
            </ul>
          </div>
          <button
            className="navbar-items-icon"
            onClick={() => setSidebar(!sidebar)}
          ></button>
        </div>
      </header>
    </>
  );
};

export default Navbar;
