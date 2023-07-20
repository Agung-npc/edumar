import React from "react";
import "./style.css";

const Hero = () => {
  return (
    <section className="hero">
      <div className="container flex">
        <div className="left ">
          <h1>Let's Level Up Your English With Us!</h1>
          <button className="margin-button">
            <a href="#card" className="primary-btn">
              Get Started
            </a>
          </button>
        </div>
        <div className="img">
          <img src="./assets/home.png" alt="" />
        </div>
      </div>
    </section>
  );
};

export default Hero;
