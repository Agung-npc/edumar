import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Card } from "react-bootstrap";
import axios from "axios";
import API from "../../../api/Api";
import "./style.css";

const Cards = () => {
  const [categories, setCategories] = useState([]);
  const navigate = useNavigate();

  const fetchCategories = async () => {
    try {
      let auth = localStorage.getItem("token");
      let { data: resp } = await axios.get(
        `${API.API_URL}/api/home/categories`,
        {
          headers: {
            Accept: "/",
            "Content-Type": "application/json",
            Authorization: "Bearer " + auth,
          },
        }
      );

      setCategories(resp.data);
    } catch (err) {}
  };

  useEffect(() => {
    fetchCategories();
  }, []);

  const handleClick = (e) => {
    navigate({
      pathname: "/quizzes",
      search: `?category_id=${e}&page=1`,
    });
  };

  return (
    <>
      <section className="about topMarign" id="card">
        <div className="heading">
          <div className="cards">
            <center>
              <h1>Choose a Course!</h1>{" "}
            </center>
            <div className="container flex d-flex gap-4 mt-5">
              <Card
                style={{ width: "18rem" }}
                className="border-0  bg-transparent"
              >
                <Card.Body className="card-body1">
                  <center>
                    <img className="card-img1" src="./assets/vocab.png" />
                  </center>
                  <Card.Title style={{ color: "white" }}>
                    <center>{categories[0]?.name}</center>
                  </Card.Title>
                  <Card.Text style={{ color: "white", textAlign: "justify" }}>
                    {categories[0]?.description}
                  </Card.Text>
                  <center>
                    <button
                      className="start1"
                      onClick={(e) => handleClick(e.target.value)}
                      value={categories[0]?.id}
                    >
                      Start
                    </button>
                  </center>
                </Card.Body>
              </Card>

              <Card
                style={{ width: "18rem" }}
                className="border-0 bg-transparent"
              >
                <Card.Body className="card-body2">
                  <center>
                    <img className="card-img2" src="./assets/grammar.png" />
                  </center>
                  <Card.Title style={{ color: "white" }}>
                    <center>{categories[1]?.name}</center>
                  </Card.Title>
                  <Card.Text style={{ color: "white", textAlign: "justify" }}>
                    {categories[1]?.description}
                  </Card.Text>
                  <center>
                    <button
                      className="start2"
                      onClick={(e) => handleClick(e.target.value)}
                      value={categories[1]?.id}
                    >
                      Start
                    </button>
                  </center>
                </Card.Body>
              </Card>

              <Card
                style={{ width: "18rem" }}
                className="border-0  bg-transparent"
              >
                <Card.Body className="card-body3">
                  <center>
                    <img className="card-img3" src="./assets/tenses.png" />
                  </center>
                  <Card.Title style={{ color: "white" }}>
                    <center>{categories[2]?.name}</center>
                  </Card.Title>
                  <Card.Text style={{ color: "white", textAlign: "justify" }}>
                    {categories[2]?.description}
                  </Card.Text>
                  <center>
                    <button
                      className="start3"
                      onClick={(e) => handleClick(e.target.value)}
                      value={categories[2]?.id}
                    >
                      Start
                    </button>
                  </center>
                </Card.Body>
              </Card>
            </div>
          </div>
        </div>
      </section>

      <div className="boxCard"></div>
    </>
  );
};

export default Cards;
