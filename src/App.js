import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <React.Fragment>
      <div class="jumbotron ">
        <h1 class="display-4">OpenShift Test</h1>
        <p class="lead">Container Information:</p>
        <hr class="my-4" />
        <table class="table">
          <tbody>
            <tr>
              <th scope="row" />
              <td>Hostname:</td>
              <td>{process.env.HOSTNAME}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </React.Fragment>
  );
}

export default App;
