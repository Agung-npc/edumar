import { Routes, Route, Navigate } from "react-router-dom";
import PrivateRoute from "./routes/PrivateRoutes";
import {
  LoginPage,
  RegisterPage,
  HomePage,
  QuizPage,
  ResultPage,
  ScoreBoardPage,
} from "./pages/switch.js";
import Navbar from "./components/Navbar/Navbar";
import Footer from "./components/Footer/Footer";
import "./App.css";

function App() {
  return (
    <>
      <Navbar />
      <Routes>
        <Route path="/login" element={<LoginPage />}></Route>
        <Route path="/register" element={<RegisterPage />}></Route>

        <Route path="" element={<PrivateRoute />}>
          <Route path="/" element={<HomePage />}></Route>
          <Route path="/quizzes" element={<QuizPage />}></Route>
          <Route path="/result" element={<ResultPage />}></Route>
          <Route path="/scores-board" element={<ScoreBoardPage />}></Route>
        </Route>

        <Route path="*" element={<Navigate to="/" />}></Route>
      </Routes>
      <Footer />
    </>
  );
}

export default App;
