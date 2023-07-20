import React from "react";
import Register from "../../components/Auth/Register";
import "./style.css";

const RegisterPage = () => {
  return (
    <div className="gradient">
      <Register />

      <div className="register-image">
        <img src="./assets/greating-image.png" height={400} width={400} />
      </div>
    </div>
  );
};

export default RegisterPage;
