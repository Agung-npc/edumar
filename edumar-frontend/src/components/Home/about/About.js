import React from "react";
import "./style.css";

const About = () => {
  return (
    <section>
      <div className="card-holder">
        <center>
          <h1>About Us</h1>
        </center>
        <p className="about" style={{ textAlign: "justify" }}>
          Edumar also known as Educational Grammar is a website-based english
          quiz where there are three types of courses namely Vocab, Grammar and
          Tenses. Each course contains 10 to 20 different questions with an
          adjusted level of difficulty. After completing one of the courses, you
          can get a gift in the form of a certificate. This website is very
          useful for practicing your English skills and of course it's free.
          This website is worked on by 6 people consisting of 4 frontends and 2
          backends.
        </p>
      </div>
    </section>
  );
};

export default About;
