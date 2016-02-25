import React from 'react';
import Relay from 'react-relay';

const App = props => {
  
  
    <div>
      <h1>{this.props.latestPost.text}</h1>
    </div>
  
  
}

export default Relay.createContainer(App, {
  fragments: {
    latestPost: () => Relay.QL`
      fragment on Post {
        text
      }
    `,
  },
});