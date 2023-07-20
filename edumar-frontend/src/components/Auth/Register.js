import React from "react";
import { Link, useNavigate } from "react-router-dom";
import axios from "axios";
import API from "../../api/Api";

const Register = () => {
  const [username, setUsername] = React.useState("");
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");

  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      let { data: res } = await axios.post(
        `${API.API_URL}/api/auth/register`,
        {
          username: username,
          email: email,
          password: password,
        },
        {
          headers: {
            Accept: "/",
            "Content-Type": "application/json",
          },
        }
      );
      if (res.code === 200) {
        navigate("/");
      }
    } catch (error) {
      alert(
        "Username / Email Sudah terdaftar, Silahkan Periksa Data Anda Kembali!"
      );
    }
  };

  return (
    <div>
      <div className="greetings-register">
        <h1> Hi, </h1>
        <h1> Welcome to</h1>
        <h1>EDUMAR!</h1>
      </div>
      <div className="container-register">
        <div className="register-box">
          <form onSubmit={handleSubmit}>
            <h2>Create new account</h2>
            <label>Name</label>
            <input
              type="text"
              placeholder="Name"
              id="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            ></input>
            <label>Email</label>
            <input
              type="text"
              placeholder="Email"
              id="email"
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
            <button type="submit" value="Register">
              Register
            </button>
            <div className="signuplink">
              <a value>Already have an account? </a>
              <Link to="/login">Login</Link>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Register;
