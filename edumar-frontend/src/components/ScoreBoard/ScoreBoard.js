import React, { useState, useEffect } from "react";
import { Table } from "react-bootstrap";
import axios from "axios";
import API from "../../api/Api";
import "./style.css";

const Board = () => {
  const [categories, setCategories] = useState([]);
  const [scoresBoard, setScoresBoard] = useState([]);

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

  const handleClick = async (e) => {
    try {
      let auth = localStorage.getItem("token");
      let { data: resp } = await axios.get(
        `${API.API_URL}/api/home/score-boards?category_id=${e.target.value}`,
        {
          headers: {
            Accept: "/",
            "Content-Type": "application/json",
            Authorization: "Bearer " + auth,
          },
        }
      );

      setScoresBoard(resp.data);
    } catch (err) {}
  };

  useEffect(() => {
    fetchCategories();
    handleClick();
  }, []);

  return (
    <div className="board">
      <h1 className="scoreboard">Scoreboard</h1>

      <div className="courses">
        {categories.map((category) => (
          <button key={category.id} onClick={handleClick} value={category.id}>
            {category.name}
          </button>
        ))}
      </div>

      <div className="container scores mt-lg-5">
        <Table striped bordered hover>
          <thead>
            <tr>
              <th>Username</th>
              <th>Score</th>
              <th>Duration</th>
            </tr>
          </thead>
          <tbody>
            {scoresBoard.map((scoreBoard, index) => {
              return scoresBoard.length !== 0 ? (
                <tr key={index}>
                  <td>{scoreBoard.username}</td>
                  <td>{scoreBoard.score}</td>
                  <td>{scoreBoard.duration}</td>
                </tr>
              ) : (
                <tr key={index}>
                  <td>{scoreBoard.duration}</td>
                </tr>
              );
            })}
          </tbody>
        </Table>
      </div>
    </div>
  );
};

export default Board;
