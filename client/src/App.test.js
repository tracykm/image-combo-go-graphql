import React from "react";
import ReactDOM from "react-dom";
import RandomImagePage from "./RandomImagePage";

it("renders without crashing", () => {
  const div = document.createElement("div");
  ReactDOM.render(<RandomImagePage />, div);
  ReactDOM.unmountComponentAtNode(div);
});
