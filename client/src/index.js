import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
import Routes from "./Routes";
import registerServiceWorker from "./registerServiceWorker";

import ApolloClient from "apollo-boost";
import { ApolloProvider } from "react-apollo";

// Pass your GraphQL endpoint to uri
const client = new ApolloClient({ uri: "http://localhost:8080/graphql" });

const ApolloApp = AppComponent => (
  <ApolloProvider client={client}>
    <AppComponent />
  </ApolloProvider>
);

ReactDOM.render(ApolloApp(Routes), document.getElementById("root"));
registerServiceWorker();
