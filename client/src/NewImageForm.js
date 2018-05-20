import React from "react";
import Button from "material-ui/Button";
import TextField from "material-ui/TextField";

import gql from "graphql-tag";
import { Mutation } from "react-apollo";

const ADD_IMG = gql`
  mutation createSrcImage($url: String!, $size: Int!, $category: String!) {
    createSrcImage(url: $url, size: $size, category: $category) {
      id
      url
      size
      category
    }
  }
`;

const categories = ["", "person", "plant", "object"];

class NewImageForm extends React.PureComponent {
  state = {
    url: "",
    size: "",
    category: ""
  };
  render() {
    const { images } = this.props;
    const state = this.state;
    return (
      <Mutation mutation={ADD_IMG}>
        {addImg => (
          <form style={{ padding: "15px" }} onSubmit={e => console.log(e)}>
            <h1>Add new</h1>
            <TextField
              id="url"
              label="url"
              margin="normal"
              value={this.state.url}
              onChange={e => this.setState({ url: e.target.value })}
            />
            <br />
            <TextField
              id="size"
              value={this.state.size}
              onChange={e => this.setState({ size: e.target.value })}
              label="size"
              type="number"
              margin="normal"
            />{" "}
            <br />
            <TextField
              id="category"
              select
              label="category"
              SelectProps={{
                native: true
              }}
              margin="normal"
              value={this.state.category}
              onChange={e => this.setState({ category: e.target.value })}
            >
              {categories.map(option => (
                <option key={option} value={option}>
                  {option}
                </option>
              ))}
            </TextField>
            <br />
            <Button
              variant="raised"
              color="primary"
              onClick={e => {
                addImg({ variables: state });
              }}
            >
              Submit
            </Button>
          </form>
        )}
      </Mutation>
    );
  }
}

export default NewImageForm;
