import React, { useState, useEffect, useRef } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import axios from "axios";
import API from "../../api/Api";
import "./style.css";

const Question = () => {
  let query = new URLSearchParams(useLocation().search);
  let categoryId = parseInt(query.get("category_id"));

  const [questions, setQuestions] = useState([]);
  const [options, setOptions] = useState([]);
  const [stopwatch, setStopwatch] = useState("00:00:00");
  const [pageQuestion, setPageQuestion] = useState(1);
  const [selected, setSelected] = useState();
  const interval = useRef(null);
  const navigate = useNavigate();

  function getTimeRemaining(endtime) {
    const total = Date.parse(endtime) - Date.parse(new Date());
    const hours = Math.floor(
      (total % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
    );
    const minutes = Math.floor((total % (1000 * 60 * 60)) / (1000 * 60));
    const seconds = Math.floor((total % (1000 * 60)) / 1000);

    return {
      total,
      hours,
      minutes,
      seconds,
    };
  }

  function startCount(endtime) {
    let { total, hours, minutes, seconds } = getTimeRemaining(endtime);
    if (total >= 0) {
      setStopwatch(
        (hours > 9 ? hours : "0" + hours) +
          ":" +
          (minutes > 9 ? minutes : "0" + minutes) +
          ":" +
          (seconds > 9 ? seconds : "0" + seconds)
      );
    } else {
      clearInterval(interval.current);
    }
  }

  useEffect(() => {
    const endtime = new Date(Date.parse(new Date()) + 3600 * 1000);
    interval.current = setInterval(() => startCount(endtime), 1000);
    return () => clearInterval(interval.current);
  }, []);

  useEffect(() => {
    const fetchQuestions = async () => {
      try {
        let auth = localStorage.getItem("token");
        const { data: res } = await axios.get(
          `${API.API_URL}/api/home/quizzes?category_id=${categoryId}&page=${pageQuestion}&limit=1`,
          {
            headers: {
              Accept: "/",
              "Content-Type": "application/json",
              Authorization: "Bearer " + auth,
            },
          }
        );

        setQuestions(res.data);

        setOptions(
          [res.data[0].correct_answer, ...res.data[0].incorrect_answers].sort(
            () => Math.random() - 0.5
          )
        );
      } catch (error) {}
    };

    fetchQuestions();
  }, [categoryId, pageQuestion]);

  const handleSelect = (index) => {
    if (selected === index && selected === questions[0]?.correct_answer)
      return "select";
    else if (selected === index && selected !== questions[0]?.correct_answer)
      return "wrong";
    else if (index === questions[0]?.correct_answer) return "select";
  };

  const handleQuit = () => {
    navigate({ pathname: "/" });
    setPageQuestion(0);
    setQuestions();
  };

  const handleCheck = (index) => {
    setSelected(index);
    let values = {
      quiz_id: questions[0].id,
      answer: index,
    };

    if (pageQuestion === 1) {
      let data = {
        answers: [values],
        category_id: 0,
        duration: "",
      };
      localStorage.setItem("answer", JSON.stringify(data));
    } else if (pageQuestion > 1) {
      let data = JSON.parse(localStorage.getItem("answer"));
      data.answers.push(values);
      localStorage.setItem("answer", JSON.stringify(data));

      if (pageQuestion === 10) {
        let timeStart = "00:59:59";
        timeStart = timeStart.split(":");

        let timeRemaining = stopwatch;
        timeRemaining = timeRemaining.split(":");

        let data = JSON.parse(localStorage.getItem("answer"));
        data.category_id = categoryId;
        data.duration = `${parseInt(
          timeStart[1] - timeRemaining[1]
        )}:${parseInt(timeStart[2] - timeRemaining[2])}`;
        localStorage.setItem("answer", JSON.stringify(data));
        console.log(localStorage.getItem("answer"));
      }
    }
  };

  const handleNext = async () => {
    if (pageQuestion > 9) {
      let auth = localStorage.getItem("token");
      let answers = localStorage.getItem("answer");

      try {
        let { data: res } = await axios.post(
          `${API.API_URL}/api/home/process-and-result`,
          answers,
          {
            headers: {
              Accept: "/",
              "Content-Type": "application/json",
              Authorization: "Bearer " + auth,
            },
          }
        );

        localStorage.setItem("result", JSON.stringify(res.data));
        navigate("/result");
      } catch (err) {}
    } else if (selected) {
      navigate({
        pathname: "/quizzes",
        search: `?category_id=${categoryId}&page=${pageQuestion + 1}`,
      });

      setPageQuestion(pageQuestion + 1);
      setSelected();
    }
  };

  return (
    <section id="question-page">
      <div className="container">
        <h1>Question</h1>
        <div className="card">
          <div className="card-body">
            <div className="row card-title mb-5 mt-3">
              <div className="quizInfo">
                <span> Soal ke - {pageQuestion} / 10</span>
                <span>{stopwatch}</span>
              </div>
            </div>

            <div className="card-text">
              {questions.map((item, index) => (
                <h5 key={index} className="mb-4">
                  {item.question}
                </h5>
              ))}
              <div className="options">
                {options.map((index) => (
                  <button
                    className={`singleOption  ${
                      selected && handleSelect(index)
                    }`}
                    key={index}
                    onClick={() => handleCheck(index)}
                    disabled={selected}
                    value={index}
                    style={{ color: "black" }}
                  >
                    {index}
                  </button>
                ))}
              </div>
            </div>
          </div>
          <div className="row card-title controls">
            <button
              type="button"
              className="btn btn-danger mr-3"
              style={{ width: 170 }}
              onClick={() => handleQuit()}
            >
              Quit
            </button>
            <button
              type="button"
              className="btn btn-warning mr-3"
              style={{ width: 170 }}
              onClick={() => handleNext()}
            >
              {pageQuestion > 9 ? "Submit" : "Next Question"}
            </button>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Question;
