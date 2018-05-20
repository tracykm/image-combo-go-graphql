import React from "react";
import { gql } from "apollo-boost";
import { Query } from "react-apollo";
import { random } from "lodash";

const OverlayedImages = ({ images }) => {
  return (
    <div
      style={{
        width: "100%",
        height: "100%",
        position: "absolute",
        top: 0,
        background: images.map(im => `url("${im.url}")`).join(","),
        backgroundBlendMode: "darken",
        backgroundPositionX: images
          .map(im => `${im.size * random(50, 150)}px`)
          .join(", "),
        backgroundPositionY: images
          .map(im => `${im.size * random(50, 150)}px`)
          .join(", "),
        backgroundSize: images
          .map(im => `${im.size * random(50, 150)}px`)
          .join(", ")
      }}
      alt={images[0].category}
    />
  );
};

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

const App = () => (
  <Query query={GET_IMAGES}>
    {({ loading, error, data }) => {
      if (loading) return <div>Loading...</div>;
      if (error) return <div>Error :(</div>;

      return (
        <div>
          <OverlayedImages images={data.SrcImagesRandom} />)}
        </div>
      );
    }}
  </Query>
);

export default App;
