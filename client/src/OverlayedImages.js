import React from "react";
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

export default OverlayedImages;
