import React from "react";
import { Link } from "react-router-dom";
import "./style.css";

const Result = () => {
  const result = JSON.parse(localStorage.getItem("result"));

  return (
    <div className="result">
      <h1 className="title">Result</h1>

      <div id="isResult">
        <div className="flex">
          <div className="text">
            <p className="isText">Correct</p>
            <span className="correctScore">
              <p>{result.correct}</p>
            </span>
          </div>
          <div className="text">
            <p className="isText">Wrong</p>
            <span className="wrongScore">
              <p>{result.wrong}</p>
            </span>
          </div>
          <div className="text">
            <p className="isText">Duration</p>
            <span className="times">
              <p>{result.duration}</p>
            </span>
          </div>
        </div>
      </div>

      <div className="toScoreboard">
        <button data-id="scoreboard">
          <Link
            to={"/scores-board"}
            style={{ textDecoration: "none", color: "black" }}
          >
            Scoreboard
          </Link>
        </button>
      </div>
    </div>
  );
};

export default Result;
