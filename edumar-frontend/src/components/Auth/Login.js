import React, { useState, useContext } from "react";
import { Link, useNavigate } from "react-router-dom";
import AuthContext from "../../store/AuthContext";
import API from "../../api/Api";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const authCtx = useContext(AuthContext);
  const navigate = useNavigate();

  if (authCtx.loggedin) {
    navigate("/");
  }

  const handleSubmit = (event) => {
    event.preventDefault();
    let auth = localStorage.getItem("token");

    let url;
    url = `${API.API_URL}/api/auth/login`;
    fetch(url, {
      method: "POST",
      body: JSON.stringify({
        email: email,
        password: password,
        returnSecureToken: true,
      }),
      headers: {
        Accept: "/",
        "Content-Type": "application/json",
        Authorization: "Bearer " + auth,
      },
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        } else {
          return res.json().then((data) => {
            let errorMassage = "Authentication failed!";
            throw new Error(errorMassage);
          });
        }
      })
      .then((data) => {
        authCtx.login(data.data.token);
      })
      .catch((err) => {
        alert("Data yang anda masukkan salah, silahkan di cek!");
      });
  };

  return (
    <div>
      <div className="greetings">
        <h1> Hallo, </h1>
        <h1> Welcome Back!</h1>
      </div>
      <div className="container-login">
        <div className="login-box">
          <form onSubmit={handleSubmit}>
            <h2>Log In</h2>
            <label>Email</label>
            <input
              type="text"
              placeholder="Email"
              id="useremail"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            ></input>

            <label>Password</label>
            <input
              type="password"
              placeholder="Password"
              id="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            ></input>

            <button type="submit" value="Log In">
              Log In
            </button>

            <div className="signuplink">
              <a value>Don't have an account? </a>
              <Link to="/register">SignUp here</Link>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Login;
