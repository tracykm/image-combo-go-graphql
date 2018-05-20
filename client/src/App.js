import React from "react";
import { gql } from "apollo-boost";
import { Query } from "react-apollo";

const Image = ({ url, size }) => {
  return <img src={url} />;
};

const GET_IMAGES = gql`
  query {
    SrcImagesRandom {
      id
      url
      category
    }
  }
`;

const App = () => (
  <Query query={GET_IMAGES}>
    {({ loading, error, data }) => {
      if (loading) return <div>Loading...</div>;
      if (error) return <div>Error :(</div>;

      return <div>{data.SrcImagesRandom.map(d => <Image {...d} />)}</div>;
    }}
  </Query>
);

export default App;
