import React from "react";
import Login from "../../components/Auth/Login";
import "./style.css";

const LoginPage = () => {
  return (
    <div className="gradient">
      <Login />

      <div className="login-image">
        <img src="./assets/greating-image.png" height={400} width={400} />
      </div>
    </div>
  );
};

export default LoginPage;
