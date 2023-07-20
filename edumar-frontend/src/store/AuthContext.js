import React, { useState } from "react";

const AuthContext = React.createContext({
  token: "",
  loggedin: false,
  login: (token) => {},
  logout: () => {},
});

export const AuthContextProvider = (props) => {
  const initialToken = localStorage.getItem("token");
  const [token, setToken] = useState(initialToken);
  const [loggedin, setLoggedin] = useState(false);

  const loginHandler = (token) => {
    setToken(token);
    setLoggedin(true);
    localStorage.setItem("token", token);
  };

  const logoutHandler = () => {
    setToken(null);
    localStorage.removeItem("token");
    localStorage.removeItem("answer");
    localStorage.removeItem("result");
  };

  const contextValue = {
    token: token,
    loggedin: loggedin,
    login: loginHandler,
    logout: logoutHandler,
  };

  return (
    <AuthContext.Provider value={contextValue}>
      {props.children}
    </AuthContext.Provider>
  );
};

export default AuthContext;
