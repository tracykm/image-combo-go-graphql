import React from "react";
import { gql } from "apollo-boost";
import { Query } from "react-apollo";
import OverlayedImages from "./OverlayedImages";

const GET_IMAGES = gql`
  query {
    SrcImagesRandom {
      id
      url
      size
      category
    }
  }
`;

const RandomImagePage = () => (
  <Query query={GET_IMAGES}>
    {({ loading, error, data }) => {
      if (loading) return <div>Loading...</div>;
      if (error) return <div>Error :(</div>;

      return (
        <div style={{ position: "relative", height: "800px" }}>
          <OverlayedImages images={data.SrcImagesRandom} />)}
        </div>
      );
    }}
  </Query>
);

export default RandomImagePage;
