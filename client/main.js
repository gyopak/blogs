import { React, ReactDOM } from 'https://unpkg.com/es-react@16.8.30';
import htm from 'https://unpkg.com/htm?module';
import Counter from './Counter.js';

const html = htm.bind(React.createElement)

ReactDOM.render(
  html`
    <h1>Look Ma! No build, no bundle, no problem</h1>
    <${Counter} class="counter" defaultCount=0 />
  `,
  document.querySelector('#app')
);
