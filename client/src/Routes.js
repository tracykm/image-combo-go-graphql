import React from "react";
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import RandomImagePage from "./RandomImagePage";
import NewImageForm from "./NewImageForm";

const Routes = () => (
  <Router>
    <div>
      <ul>
        <li>
          <Link to="/">Home</Link>
        </li>
        <li>
          <Link to="/random">Random</Link>
        </li>
        <li>
          <Link to="/new">New</Link>
        </li>
      </ul>

      <hr />

      <Route exact path="/" component={RandomImagePage} />
      <Route path="/random" component={RandomImagePage} />
      <Route path="/new" component={NewImageForm} />
    </div>
  </Router>
);

export default Routes;
