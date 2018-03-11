package main

var playerCSS = `
.asciinema-player-wrapper {
  position: relative;
  text-align: center;
  outline: none;
}
.asciinema-player-wrapper .title-bar {
  display: none;
  top: -78px;
  transition: top 0.15s linear;
  position: absolute;
  left: 0;
  right: 0;
  box-sizing: content-box;
  font-size: 20px;
  line-height: 1em;
  padding: 15px;
  font-family: sans-serif;
  color: white;
  background-color: rgba(0, 0, 0, 0.8);
}
.asciinema-player-wrapper .title-bar img {
  vertical-align: middle;
  height: 48px;
  margin-right: 16px;
}
.asciinema-player-wrapper .title-bar a {
  color: white;
  text-decoration: underline;
}
.asciinema-player-wrapper .title-bar a:hover {
  text-decoration: none;
}
.asciinema-player-wrapper:fullscreen {
  background-color: #000;
  width: 100%;
  height: 100%;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  -webkit-justify-content: center;
  justify-content: center;
  -webkit-align-items: center;
  align-items: center;
}
.asciinema-player-wrapper:fullscreen .asciinema-player {
  position: static;
}
.asciinema-player-wrapper:fullscreen .title-bar {
  display: initial;
}
.asciinema-player-wrapper:fullscreen.hud .title-bar {
  top: 0;
}
.asciinema-player-wrapper:-webkit-full-screen {
  background-color: #000;
  width: 100%;
  height: 100%;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  -webkit-justify-content: center;
  justify-content: center;
  -webkit-align-items: center;
  align-items: center;
}
.asciinema-player-wrapper:-webkit-full-screen .asciinema-player {
  position: static;
}
.asciinema-player-wrapper:-webkit-full-screen .title-bar {
  display: initial;
}
.asciinema-player-wrapper:-webkit-full-screen.hud .title-bar {
  top: 0;
}
.asciinema-player-wrapper:-moz-full-screen {
  background-color: #000;
  width: 100%;
  height: 100%;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  -webkit-justify-content: center;
  justify-content: center;
  -webkit-align-items: center;
  align-items: center;
}
.asciinema-player-wrapper:-moz-full-screen .asciinema-player {
  position: static;
}
.asciinema-player-wrapper:-moz-full-screen .title-bar {
  display: initial;
}
.asciinema-player-wrapper:-moz-full-screen.hud .title-bar {
  top: 0;
}
.asciinema-player-wrapper:-ms-fullscreen {
  background-color: #000;
  width: 100%;
  height: 100%;
  display: -webkit-flex;
  display: -ms-flexbox;
  display: flex;
  -webkit-justify-content: center;
  justify-content: center;
  -webkit-align-items: center;
  align-items: center;
}
.asciinema-player-wrapper:-ms-fullscreen .asciinema-player {
  position: static;
}
.asciinema-player-wrapper:-ms-fullscreen .title-bar {
  display: initial;
}
.asciinema-player-wrapper:-ms-fullscreen.hud .title-bar {
  top: 0;
}
.asciinema-player-wrapper .asciinema-player {
  text-align: left;
  display: inline-block;
  padding: 0px;
  position: relative;
  box-sizing: content-box;
  -moz-box-sizing: content-box;
  -webkit-box-sizing: content-box;
  overflow: hidden;
  max-width: 100%;
}
.asciinema-terminal {
  box-sizing: content-box;
  -moz-box-sizing: content-box;
  -webkit-box-sizing: content-box;
  overflow: hidden;
  padding: 0;
  margin: 0px;
  display: block;
  white-space: pre;
  border: 0;
  word-wrap: normal;
  word-break: normal;
  border-radius: 0;
  border-style: solid;
  cursor: text;
  border-width: 0.5em;
  font-family: Consolas, Menlo, 'Bitstream Vera Sans Mono', monospace, 'Powerline Symbols';
  line-height: 1.3333333333em;
}
.asciinema-terminal .line {
  letter-spacing: normal;
  overflow: hidden;
  height: 1.3333333333em;
}
.asciinema-terminal .line span {
  padding: 0;
  display: inline-block;
  height: 1.3333333333em;
}
.asciinema-terminal .line {
  display: block;
  width: 200%;
}
.asciinema-terminal .bright {
  font-weight: bold;
}
.asciinema-terminal .underline {
  text-decoration: underline;
}
.asciinema-terminal .italic {
  font-style: italic;
}
.asciinema-terminal.font-small {
  font-size: 12px;
}
.asciinema-terminal.font-medium {
  font-size: 18px;
}
.asciinema-terminal.font-big {
  font-size: 24px;
}
.asciinema-player .control-bar {
  width: 100%;
  height: 32px;
  background: rgba(0, 0, 0, 0.8);
  /* no gradient fallback */
  background: -moz-linear-gradient(top, rgba(0, 0, 0, 0.5) 0%, #000000 25%, #000000 100%);
  /* FF3.6-15 */
  background: -webkit-linear-gradient(top, rgba(0, 0, 0, 0.5) 0%, #000000 25%, #000000 100%);
  /* Chrome10-25,Safari5.1-6 */
  background: linear-gradient(to bottom, rgba(0, 0, 0, 0.5) 0%, #000000 25%, #000000 100%);
  /* W3C, IE10+, FF16+, Chrome26+, Opera12+, Safari7+ */
  color: #bbbbbb;
  box-sizing: content-box;
  line-height: 1;
  position: absolute;
  bottom: -35px;
  left: 0;
  transition: bottom 0.15s linear;
}
.asciinema-player .control-bar * {
  box-sizing: inherit;
  font-size: 0;
}
.asciinema-player .control-bar svg.icon path {
  fill: #bbbbbb;
}
.asciinema-player .control-bar .playback-button {
  display: block;
  float: left;
  cursor: pointer;
  height: 12px;
  width: 12px;
  padding: 10px;
}
.asciinema-player .control-bar .playback-button svg {
  height: 12px;
  width: 12px;
}
.asciinema-player .control-bar .timer {
  display: block;
  float: left;
  width: 50px;
  height: 100%;
  text-align: center;
  font-family: Helvetica, Arial, sans-serif;
  font-size: 11px;
  font-weight: bold;
  line-height: 32px;
  cursor: default;
}
.asciinema-player .control-bar .timer span {
  display: inline-block;
  font-size: inherit;
}
.asciinema-player .control-bar .timer .time-remaining {
  display: none;
}
.asciinema-player .control-bar .timer:hover .time-elapsed {
  display: none;
}
.asciinema-player .control-bar .timer:hover .time-remaining {
  display: inline;
}
.asciinema-player .control-bar .progressbar {
  display: block;
  overflow: hidden;
  height: 100%;
  padding: 0 10px;
}
.asciinema-player .control-bar .progressbar .bar {
  display: block;
  cursor: pointer;
  height: 100%;
  padding-top: 15px;
  font-size: 0;
}
.asciinema-player .control-bar .progressbar .bar .gutter {
  display: block;
  height: 3px;
  background-color: #333;
}
.asciinema-player .control-bar .progressbar .bar .gutter span {
  display: inline-block;
  height: 100%;
  background-color: #bbbbbb;
  border-radius: 3px;
}
.asciinema-player .control-bar.live .progressbar .bar {
  cursor: default;
}
.asciinema-player .control-bar .fullscreen-button {
  display: block;
  float: right;
  width: 14px;
  height: 14px;
  padding: 9px;
  cursor: pointer;
}
.asciinema-player .control-bar .fullscreen-button svg {
  width: 14px;
  height: 14px;
}
.asciinema-player .control-bar .fullscreen-button svg:first-child {
  display: inline;
}
.asciinema-player .control-bar .fullscreen-button svg:last-child {
  display: none;
}
.asciinema-player-wrapper.hud .control-bar {
  bottom: 0px;
}
.asciinema-player-wrapper:fullscreen .fullscreen-button svg:first-child {
  display: none;
}
.asciinema-player-wrapper:fullscreen .fullscreen-button svg:last-child {
  display: inline;
}
.asciinema-player-wrapper:-webkit-full-screen .fullscreen-button svg:first-child {
  display: none;
}
.asciinema-player-wrapper:-webkit-full-screen .fullscreen-button svg:last-child {
  display: inline;
}
.asciinema-player-wrapper:-moz-full-screen .fullscreen-button svg:first-child {
  display: none;
}
.asciinema-player-wrapper:-moz-full-screen .fullscreen-button svg:last-child {
  display: inline;
}
.asciinema-player-wrapper:-ms-fullscreen .fullscreen-button svg:first-child {
  display: none;
}
.asciinema-player-wrapper:-ms-fullscreen .fullscreen-button svg:last-child {
  display: inline;
}
.asciinema-player .loading {
  z-index: 10;
  background-repeat: no-repeat;
  background-position: center;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 32px;
  background-color: rgba(0, 0, 0, 0.5);
}
.asciinema-player .start-prompt {
  z-index: 10;
  background-repeat: no-repeat;
  background-position: center;
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 32px;
  z-index: 20;
  cursor: pointer;
}
.asciinema-player .start-prompt .play-button {
  font-size: 0px;
}
.asciinema-player .start-prompt .play-button {
  position: absolute;
  left: 0;
  top: 0;
  right: 0;
  bottom: 0;
  text-align: center;
  color: white;
  display: table;
  width: 100%;
  height: 100%;
}
.asciinema-player .start-prompt .play-button div {
  vertical-align: middle;
  display: table-cell;
}
.asciinema-player .start-prompt .play-button div span {
  width: 96px;
  height: 96px;
  display: inline-block;
}
@-webkit-keyframes expand {
  0% {
    -webkit-transform: scale(0);
  }
  50% {
    -webkit-transform: scale(1);
  }
  100% {
    z-index: 1;
  }
}
@-moz-keyframes expand {
  0% {
    -moz-transform: scale(0);
  }
  50% {
    -moz-transform: scale(1);
  }
  100% {
    z-index: 1;
  }
}
@-o-keyframes expand {
  0% {
    -o-transform: scale(0);
  }
  50% {
    -o-transform: scale(1);
  }
  100% {
    z-index: 1;
  }
}
@keyframes expand {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1);
  }
  100% {
    z-index: 1;
  }
}
.loader {
  position: absolute;
  left: 50%;
  top: 50%;
  margin: -20px 0 0 -20px;
  background-color: white;
  border-radius: 50%;
  box-shadow: 0 0 0 6.66667px #141414;
  width: 40px;
  height: 40px;
}
.loader:before,
.loader:after {
  content: "";
  position: absolute;
  left: 50%;
  top: 50%;
  display: block;
  margin: -21px 0 0 -21px;
  border-radius: 50%;
  z-index: 2;
  width: 42px;
  height: 42px;
}
.loader:before {
  background-color: #141414;
  -webkit-animation: expand 1.6s linear infinite both;
  -moz-animation: expand 1.6s linear infinite both;
  animation: expand 1.6s linear infinite both;
}
.loader:after {
  background-color: white;
  -webkit-animation: expand 1.6s linear 0.8s infinite both;
  -moz-animation: expand 1.6s linear 0.8s infinite both;
  animation: expand 1.6s linear 0.8s infinite both;
}
.asciinema-terminal .fg-16 {
  color: #000000;
}
.asciinema-terminal .bg-16 {
  background-color: #000000;
}
.asciinema-terminal .fg-17 {
  color: #00005f;
}
.asciinema-terminal .bg-17 {
  background-color: #00005f;
}
.asciinema-terminal .fg-18 {
  color: #000087;
}
.asciinema-terminal .bg-18 {
  background-color: #000087;
}
.asciinema-terminal .fg-19 {
  color: #0000af;
}
.asciinema-terminal .bg-19 {
  background-color: #0000af;
}
.asciinema-terminal .fg-20 {
  color: #0000d7;
}
.asciinema-terminal .bg-20 {
  background-color: #0000d7;
}
.asciinema-terminal .fg-21 {
  color: #0000ff;
}
.asciinema-terminal .bg-21 {
  background-color: #0000ff;
}
.asciinema-terminal .fg-22 {
  color: #005f00;
}
.asciinema-terminal .bg-22 {
  background-color: #005f00;
}
.asciinema-terminal .fg-23 {
  color: #005f5f;
}
.asciinema-terminal .bg-23 {
  background-color: #005f5f;
}
.asciinema-terminal .fg-24 {
  color: #005f87;
}
.asciinema-terminal .bg-24 {
  background-color: #005f87;
}
.asciinema-terminal .fg-25 {
  color: #005faf;
}
.asciinema-terminal .bg-25 {
  background-color: #005faf;
}
.asciinema-terminal .fg-26 {
  color: #005fd7;
}
.asciinema-terminal .bg-26 {
  background-color: #005fd7;
}
.asciinema-terminal .fg-27 {
  color: #005fff;
}
.asciinema-terminal .bg-27 {
  background-color: #005fff;
}
.asciinema-terminal .fg-28 {
  color: #008700;
}
.asciinema-terminal .bg-28 {
  background-color: #008700;
}
.asciinema-terminal .fg-29 {
  color: #00875f;
}
.asciinema-terminal .bg-29 {
  background-color: #00875f;
}
.asciinema-terminal .fg-30 {
  color: #008787;
}
.asciinema-terminal .bg-30 {
  background-color: #008787;
}
.asciinema-terminal .fg-31 {
  color: #0087af;
}
.asciinema-terminal .bg-31 {
  background-color: #0087af;
}
.asciinema-terminal .fg-32 {
  color: #0087d7;
}
.asciinema-terminal .bg-32 {
  background-color: #0087d7;
}
.asciinema-terminal .fg-33 {
  color: #0087ff;
}
.asciinema-terminal .bg-33 {
  background-color: #0087ff;
}
.asciinema-terminal .fg-34 {
  color: #00af00;
}
.asciinema-terminal .bg-34 {
  background-color: #00af00;
}
.asciinema-terminal .fg-35 {
  color: #00af5f;
}
.asciinema-terminal .bg-35 {
  background-color: #00af5f;
}
.asciinema-terminal .fg-36 {
  color: #00af87;
}
.asciinema-terminal .bg-36 {
  background-color: #00af87;
}
.asciinema-terminal .fg-37 {
  color: #00afaf;
}
.asciinema-terminal .bg-37 {
  background-color: #00afaf;
}
.asciinema-terminal .fg-38 {
  color: #00afd7;
}
.asciinema-terminal .bg-38 {
  background-color: #00afd7;
}
.asciinema-terminal .fg-39 {
  color: #00afff;
}
.asciinema-terminal .bg-39 {
  background-color: #00afff;
}
.asciinema-terminal .fg-40 {
  color: #00d700;
}
.asciinema-terminal .bg-40 {
  background-color: #00d700;
}
.asciinema-terminal .fg-41 {
  color: #00d75f;
}
.asciinema-terminal .bg-41 {
  background-color: #00d75f;
}
.asciinema-terminal .fg-42 {
  color: #00d787;
}
.asciinema-terminal .bg-42 {
  background-color: #00d787;
}
.asciinema-terminal .fg-43 {
  color: #00d7af;
}
.asciinema-terminal .bg-43 {
  background-color: #00d7af;
}
.asciinema-terminal .fg-44 {
  color: #00d7d7;
}
.asciinema-terminal .bg-44 {
  background-color: #00d7d7;
}
.asciinema-terminal .fg-45 {
  color: #00d7ff;
}
.asciinema-terminal .bg-45 {
  background-color: #00d7ff;
}
.asciinema-terminal .fg-46 {
  color: #00ff00;
}
.asciinema-terminal .bg-46 {
  background-color: #00ff00;
}
.asciinema-terminal .fg-47 {
  color: #00ff5f;
}
.asciinema-terminal .bg-47 {
  background-color: #00ff5f;
}
.asciinema-terminal .fg-48 {
  color: #00ff87;
}
.asciinema-terminal .bg-48 {
  background-color: #00ff87;
}
.asciinema-terminal .fg-49 {
  color: #00ffaf;
}
.asciinema-terminal .bg-49 {
  background-color: #00ffaf;
}
.asciinema-terminal .fg-50 {
  color: #00ffd7;
}
.asciinema-terminal .bg-50 {
  background-color: #00ffd7;
}
.asciinema-terminal .fg-51 {
  color: #00ffff;
}
.asciinema-terminal .bg-51 {
  background-color: #00ffff;
}
.asciinema-terminal .fg-52 {
  color: #5f0000;
}
.asciinema-terminal .bg-52 {
  background-color: #5f0000;
}
.asciinema-terminal .fg-53 {
  color: #5f005f;
}
.asciinema-terminal .bg-53 {
  background-color: #5f005f;
}
.asciinema-terminal .fg-54 {
  color: #5f0087;
}
.asciinema-terminal .bg-54 {
  background-color: #5f0087;
}
.asciinema-terminal .fg-55 {
  color: #5f00af;
}
.asciinema-terminal .bg-55 {
  background-color: #5f00af;
}
.asciinema-terminal .fg-56 {
  color: #5f00d7;
}
.asciinema-terminal .bg-56 {
  background-color: #5f00d7;
}
.asciinema-terminal .fg-57 {
  color: #5f00ff;
}
.asciinema-terminal .bg-57 {
  background-color: #5f00ff;
}
.asciinema-terminal .fg-58 {
  color: #5f5f00;
}
.asciinema-terminal .bg-58 {
  background-color: #5f5f00;
}
.asciinema-terminal .fg-59 {
  color: #5f5f5f;
}
.asciinema-terminal .bg-59 {
  background-color: #5f5f5f;
}
.asciinema-terminal .fg-60 {
  color: #5f5f87;
}
.asciinema-terminal .bg-60 {
  background-color: #5f5f87;
}
.asciinema-terminal .fg-61 {
  color: #5f5faf;
}
.asciinema-terminal .bg-61 {
  background-color: #5f5faf;
}
.asciinema-terminal .fg-62 {
  color: #5f5fd7;
}
.asciinema-terminal .bg-62 {
  background-color: #5f5fd7;
}
.asciinema-terminal .fg-63 {
  color: #5f5fff;
}
.asciinema-terminal .bg-63 {
  background-color: #5f5fff;
}
.asciinema-terminal .fg-64 {
  color: #5f8700;
}
.asciinema-terminal .bg-64 {
  background-color: #5f8700;
}
.asciinema-terminal .fg-65 {
  color: #5f875f;
}
.asciinema-terminal .bg-65 {
  background-color: #5f875f;
}
.asciinema-terminal .fg-66 {
  color: #5f8787;
}
.asciinema-terminal .bg-66 {
  background-color: #5f8787;
}
.asciinema-terminal .fg-67 {
  color: #5f87af;
}
.asciinema-terminal .bg-67 {
  background-color: #5f87af;
}
.asciinema-terminal .fg-68 {
  color: #5f87d7;
}
.asciinema-terminal .bg-68 {
  background-color: #5f87d7;
}
.asciinema-terminal .fg-69 {
  color: #5f87ff;
}
.asciinema-terminal .bg-69 {
  background-color: #5f87ff;
}
.asciinema-terminal .fg-70 {
  color: #5faf00;
}
.asciinema-terminal .bg-70 {
  background-color: #5faf00;
}
.asciinema-terminal .fg-71 {
  color: #5faf5f;
}
.asciinema-terminal .bg-71 {
  background-color: #5faf5f;
}
.asciinema-terminal .fg-72 {
  color: #5faf87;
}
.asciinema-terminal .bg-72 {
  background-color: #5faf87;
}
.asciinema-terminal .fg-73 {
  color: #5fafaf;
}
.asciinema-terminal .bg-73 {
  background-color: #5fafaf;
}
.asciinema-terminal .fg-74 {
  color: #5fafd7;
}
.asciinema-terminal .bg-74 {
  background-color: #5fafd7;
}
.asciinema-terminal .fg-75 {
  color: #5fafff;
}
.asciinema-terminal .bg-75 {
  background-color: #5fafff;
}
.asciinema-terminal .fg-76 {
  color: #5fd700;
}
.asciinema-terminal .bg-76 {
  background-color: #5fd700;
}
.asciinema-terminal .fg-77 {
  color: #5fd75f;
}
.asciinema-terminal .bg-77 {
  background-color: #5fd75f;
}
.asciinema-terminal .fg-78 {
  color: #5fd787;
}
.asciinema-terminal .bg-78 {
  background-color: #5fd787;
}
.asciinema-terminal .fg-79 {
  color: #5fd7af;
}
.asciinema-terminal .bg-79 {
  background-color: #5fd7af;
}
.asciinema-terminal .fg-80 {
  color: #5fd7d7;
}
.asciinema-terminal .bg-80 {
  background-color: #5fd7d7;
}
.asciinema-terminal .fg-81 {
  color: #5fd7ff;
}
.asciinema-terminal .bg-81 {
  background-color: #5fd7ff;
}
.asciinema-terminal .fg-82 {
  color: #5fff00;
}
.asciinema-terminal .bg-82 {
  background-color: #5fff00;
}
.asciinema-terminal .fg-83 {
  color: #5fff5f;
}
.asciinema-terminal .bg-83 {
  background-color: #5fff5f;
}
.asciinema-terminal .fg-84 {
  color: #5fff87;
}
.asciinema-terminal .bg-84 {
  background-color: #5fff87;
}
.asciinema-terminal .fg-85 {
  color: #5fffaf;
}
.asciinema-terminal .bg-85 {
  background-color: #5fffaf;
}
.asciinema-terminal .fg-86 {
  color: #5fffd7;
}
.asciinema-terminal .bg-86 {
  background-color: #5fffd7;
}
.asciinema-terminal .fg-87 {
  color: #5fffff;
}
.asciinema-terminal .bg-87 {
  background-color: #5fffff;
}
.asciinema-terminal .fg-88 {
  color: #870000;
}
.asciinema-terminal .bg-88 {
  background-color: #870000;
}
.asciinema-terminal .fg-89 {
  color: #87005f;
}
.asciinema-terminal .bg-89 {
  background-color: #87005f;
}
.asciinema-terminal .fg-90 {
  color: #870087;
}
.asciinema-terminal .bg-90 {
  background-color: #870087;
}
.asciinema-terminal .fg-91 {
  color: #8700af;
}
.asciinema-terminal .bg-91 {
  background-color: #8700af;
}
.asciinema-terminal .fg-92 {
  color: #8700d7;
}
.asciinema-terminal .bg-92 {
  background-color: #8700d7;
}
.asciinema-terminal .fg-93 {
  color: #8700ff;
}
.asciinema-terminal .bg-93 {
  background-color: #8700ff;
}
.asciinema-terminal .fg-94 {
  color: #875f00;
}
.asciinema-terminal .bg-94 {
  background-color: #875f00;
}
.asciinema-terminal .fg-95 {
  color: #875f5f;
}
.asciinema-terminal .bg-95 {
  background-color: #875f5f;
}
.asciinema-terminal .fg-96 {
  color: #875f87;
}
.asciinema-terminal .bg-96 {
  background-color: #875f87;
}
.asciinema-terminal .fg-97 {
  color: #875faf;
}
.asciinema-terminal .bg-97 {
  background-color: #875faf;
}
.asciinema-terminal .fg-98 {
  color: #875fd7;
}
.asciinema-terminal .bg-98 {
  background-color: #875fd7;
}
.asciinema-terminal .fg-99 {
  color: #875fff;
}
.asciinema-terminal .bg-99 {
  background-color: #875fff;
}
.asciinema-terminal .fg-100 {
  color: #878700;
}
.asciinema-terminal .bg-100 {
  background-color: #878700;
}
.asciinema-terminal .fg-101 {
  color: #87875f;
}
.asciinema-terminal .bg-101 {
  background-color: #87875f;
}
.asciinema-terminal .fg-102 {
  color: #878787;
}
.asciinema-terminal .bg-102 {
  background-color: #878787;
}
.asciinema-terminal .fg-103 {
  color: #8787af;
}
.asciinema-terminal .bg-103 {
  background-color: #8787af;
}
.asciinema-terminal .fg-104 {
  color: #8787d7;
}
.asciinema-terminal .bg-104 {
  background-color: #8787d7;
}
.asciinema-terminal .fg-105 {
  color: #8787ff;
}
.asciinema-terminal .bg-105 {
  background-color: #8787ff;
}
.asciinema-terminal .fg-106 {
  color: #87af00;
}
.asciinema-terminal .bg-106 {
  background-color: #87af00;
}
.asciinema-terminal .fg-107 {
  color: #87af5f;
}
.asciinema-terminal .bg-107 {
  background-color: #87af5f;
}
.asciinema-terminal .fg-108 {
  color: #87af87;
}
.asciinema-terminal .bg-108 {
  background-color: #87af87;
}
.asciinema-terminal .fg-109 {
  color: #87afaf;
}
.asciinema-terminal .bg-109 {
  background-color: #87afaf;
}
.asciinema-terminal .fg-110 {
  color: #87afd7;
}
.asciinema-terminal .bg-110 {
  background-color: #87afd7;
}
.asciinema-terminal .fg-111 {
  color: #87afff;
}
.asciinema-terminal .bg-111 {
  background-color: #87afff;
}
.asciinema-terminal .fg-112 {
  color: #87d700;
}
.asciinema-terminal .bg-112 {
  background-color: #87d700;
}
.asciinema-terminal .fg-113 {
  color: #87d75f;
}
.asciinema-terminal .bg-113 {
  background-color: #87d75f;
}
.asciinema-terminal .fg-114 {
  color: #87d787;
}
.asciinema-terminal .bg-114 {
  background-color: #87d787;
}
.asciinema-terminal .fg-115 {
  color: #87d7af;
}
.asciinema-terminal .bg-115 {
  background-color: #87d7af;
}
.asciinema-terminal .fg-116 {
  color: #87d7d7;
}
.asciinema-terminal .bg-116 {
  background-color: #87d7d7;
}
.asciinema-terminal .fg-117 {
  color: #87d7ff;
}
.asciinema-terminal .bg-117 {
  background-color: #87d7ff;
}
.asciinema-terminal .fg-118 {
  color: #87ff00;
}
.asciinema-terminal .bg-118 {
  background-color: #87ff00;
}
.asciinema-terminal .fg-119 {
  color: #87ff5f;
}
.asciinema-terminal .bg-119 {
  background-color: #87ff5f;
}
.asciinema-terminal .fg-120 {
  color: #87ff87;
}
.asciinema-terminal .bg-120 {
  background-color: #87ff87;
}
.asciinema-terminal .fg-121 {
  color: #87ffaf;
}
.asciinema-terminal .bg-121 {
  background-color: #87ffaf;
}
.asciinema-terminal .fg-122 {
  color: #87ffd7;
}
.asciinema-terminal .bg-122 {
  background-color: #87ffd7;
}
.asciinema-terminal .fg-123 {
  color: #87ffff;
}
.asciinema-terminal .bg-123 {
  background-color: #87ffff;
}
.asciinema-terminal .fg-124 {
  color: #af0000;
}
.asciinema-terminal .bg-124 {
  background-color: #af0000;
}
.asciinema-terminal .fg-125 {
  color: #af005f;
}
.asciinema-terminal .bg-125 {
  background-color: #af005f;
}
.asciinema-terminal .fg-126 {
  color: #af0087;
}
.asciinema-terminal .bg-126 {
  background-color: #af0087;
}
.asciinema-terminal .fg-127 {
  color: #af00af;
}
.asciinema-terminal .bg-127 {
  background-color: #af00af;
}
.asciinema-terminal .fg-128 {
  color: #af00d7;
}
.asciinema-terminal .bg-128 {
  background-color: #af00d7;
}
.asciinema-terminal .fg-129 {
  color: #af00ff;
}
.asciinema-terminal .bg-129 {
  background-color: #af00ff;
}
.asciinema-terminal .fg-130 {
  color: #af5f00;
}
.asciinema-terminal .bg-130 {
  background-color: #af5f00;
}
.asciinema-terminal .fg-131 {
  color: #af5f5f;
}
.asciinema-terminal .bg-131 {
  background-color: #af5f5f;
}
.asciinema-terminal .fg-132 {
  color: #af5f87;
}
.asciinema-terminal .bg-132 {
  background-color: #af5f87;
}
.asciinema-terminal .fg-133 {
  color: #af5faf;
}
.asciinema-terminal .bg-133 {
  background-color: #af5faf;
}
.asciinema-terminal .fg-134 {
  color: #af5fd7;
}
.asciinema-terminal .bg-134 {
  background-color: #af5fd7;
}
.asciinema-terminal .fg-135 {
  color: #af5fff;
}
.asciinema-terminal .bg-135 {
  background-color: #af5fff;
}
.asciinema-terminal .fg-136 {
  color: #af8700;
}
.asciinema-terminal .bg-136 {
  background-color: #af8700;
}
.asciinema-terminal .fg-137 {
  color: #af875f;
}
.asciinema-terminal .bg-137 {
  background-color: #af875f;
}
.asciinema-terminal .fg-138 {
  color: #af8787;
}
.asciinema-terminal .bg-138 {
  background-color: #af8787;
}
.asciinema-terminal .fg-139 {
  color: #af87af;
}
.asciinema-terminal .bg-139 {
  background-color: #af87af;
}
.asciinema-terminal .fg-140 {
  color: #af87d7;
}
.asciinema-terminal .bg-140 {
  background-color: #af87d7;
}
.asciinema-terminal .fg-141 {
  color: #af87ff;
}
.asciinema-terminal .bg-141 {
  background-color: #af87ff;
}
.asciinema-terminal .fg-142 {
  color: #afaf00;
}
.asciinema-terminal .bg-142 {
  background-color: #afaf00;
}
.asciinema-terminal .fg-143 {
  color: #afaf5f;
}
.asciinema-terminal .bg-143 {
  background-color: #afaf5f;
}
.asciinema-terminal .fg-144 {
  color: #afaf87;
}
.asciinema-terminal .bg-144 {
  background-color: #afaf87;
}
.asciinema-terminal .fg-145 {
  color: #afafaf;
}
.asciinema-terminal .bg-145 {
  background-color: #afafaf;
}
.asciinema-terminal .fg-146 {
  color: #afafd7;
}
.asciinema-terminal .bg-146 {
  background-color: #afafd7;
}
.asciinema-terminal .fg-147 {
  color: #afafff;
}
.asciinema-terminal .bg-147 {
  background-color: #afafff;
}
.asciinema-terminal .fg-148 {
  color: #afd700;
}
.asciinema-terminal .bg-148 {
  background-color: #afd700;
}
.asciinema-terminal .fg-149 {
  color: #afd75f;
}
.asciinema-terminal .bg-149 {
  background-color: #afd75f;
}
.asciinema-terminal .fg-150 {
  color: #afd787;
}
.asciinema-terminal .bg-150 {
  background-color: #afd787;
}
.asciinema-terminal .fg-151 {
  color: #afd7af;
}
.asciinema-terminal .bg-151 {
  background-color: #afd7af;
}
.asciinema-terminal .fg-152 {
  color: #afd7d7;
}
.asciinema-terminal .bg-152 {
  background-color: #afd7d7;
}
.asciinema-terminal .fg-153 {
  color: #afd7ff;
}
.asciinema-terminal .bg-153 {
  background-color: #afd7ff;
}
.asciinema-terminal .fg-154 {
  color: #afff00;
}
.asciinema-terminal .bg-154 {
  background-color: #afff00;
}
.asciinema-terminal .fg-155 {
  color: #afff5f;
}
.asciinema-terminal .bg-155 {
  background-color: #afff5f;
}
.asciinema-terminal .fg-156 {
  color: #afff87;
}
.asciinema-terminal .bg-156 {
  background-color: #afff87;
}
.asciinema-terminal .fg-157 {
  color: #afffaf;
}
.asciinema-terminal .bg-157 {
  background-color: #afffaf;
}
.asciinema-terminal .fg-158 {
  color: #afffd7;
}
.asciinema-terminal .bg-158 {
  background-color: #afffd7;
}
.asciinema-terminal .fg-159 {
  color: #afffff;
}
.asciinema-terminal .bg-159 {
  background-color: #afffff;
}
.asciinema-terminal .fg-160 {
  color: #d70000;
}
.asciinema-terminal .bg-160 {
  background-color: #d70000;
}
.asciinema-terminal .fg-161 {
  color: #d7005f;
}
.asciinema-terminal .bg-161 {
  background-color: #d7005f;
}
.asciinema-terminal .fg-162 {
  color: #d70087;
}
.asciinema-terminal .bg-162 {
  background-color: #d70087;
}
.asciinema-terminal .fg-163 {
  color: #d700af;
}
.asciinema-terminal .bg-163 {
  background-color: #d700af;
}
.asciinema-terminal .fg-164 {
  color: #d700d7;
}
.asciinema-terminal .bg-164 {
  background-color: #d700d7;
}
.asciinema-terminal .fg-165 {
  color: #d700ff;
}
.asciinema-terminal .bg-165 {
  background-color: #d700ff;
}
.asciinema-terminal .fg-166 {
  color: #d75f00;
}
.asciinema-terminal .bg-166 {
  background-color: #d75f00;
}
.asciinema-terminal .fg-167 {
  color: #d75f5f;
}
.asciinema-terminal .bg-167 {
  background-color: #d75f5f;
}
.asciinema-terminal .fg-168 {
  color: #d75f87;
}
.asciinema-terminal .bg-168 {
  background-color: #d75f87;
}
.asciinema-terminal .fg-169 {
  color: #d75faf;
}
.asciinema-terminal .bg-169 {
  background-color: #d75faf;
}
.asciinema-terminal .fg-170 {
  color: #d75fd7;
}
.asciinema-terminal .bg-170 {
  background-color: #d75fd7;
}
.asciinema-terminal .fg-171 {
  color: #d75fff;
}
.asciinema-terminal .bg-171 {
  background-color: #d75fff;
}
.asciinema-terminal .fg-172 {
  color: #d78700;
}
.asciinema-terminal .bg-172 {
  background-color: #d78700;
}
.asciinema-terminal .fg-173 {
  color: #d7875f;
}
.asciinema-terminal .bg-173 {
  background-color: #d7875f;
}
.asciinema-terminal .fg-174 {
  color: #d78787;
}
.asciinema-terminal .bg-174 {
  background-color: #d78787;
}
.asciinema-terminal .fg-175 {
  color: #d787af;
}
.asciinema-terminal .bg-175 {
  background-color: #d787af;
}
.asciinema-terminal .fg-176 {
  color: #d787d7;
}
.asciinema-terminal .bg-176 {
  background-color: #d787d7;
}
.asciinema-terminal .fg-177 {
  color: #d787ff;
}
.asciinema-terminal .bg-177 {
  background-color: #d787ff;
}
.asciinema-terminal .fg-178 {
  color: #d7af00;
}
.asciinema-terminal .bg-178 {
  background-color: #d7af00;
}
.asciinema-terminal .fg-179 {
  color: #d7af5f;
}
.asciinema-terminal .bg-179 {
  background-color: #d7af5f;
}
.asciinema-terminal .fg-180 {
  color: #d7af87;
}
.asciinema-terminal .bg-180 {
  background-color: #d7af87;
}
.asciinema-terminal .fg-181 {
  color: #d7afaf;
}
.asciinema-terminal .bg-181 {
  background-color: #d7afaf;
}
.asciinema-terminal .fg-182 {
  color: #d7afd7;
}
.asciinema-terminal .bg-182 {
  background-color: #d7afd7;
}
.asciinema-terminal .fg-183 {
  color: #d7afff;
}
.asciinema-terminal .bg-183 {
  background-color: #d7afff;
}
.asciinema-terminal .fg-184 {
  color: #d7d700;
}
.asciinema-terminal .bg-184 {
  background-color: #d7d700;
}
.asciinema-terminal .fg-185 {
  color: #d7d75f;
}
.asciinema-terminal .bg-185 {
  background-color: #d7d75f;
}
.asciinema-terminal .fg-186 {
  color: #d7d787;
}
.asciinema-terminal .bg-186 {
  background-color: #d7d787;
}
.asciinema-terminal .fg-187 {
  color: #d7d7af;
}
.asciinema-terminal .bg-187 {
  background-color: #d7d7af;
}
.asciinema-terminal .fg-188 {
  color: #d7d7d7;
}
.asciinema-terminal .bg-188 {
  background-color: #d7d7d7;
}
.asciinema-terminal .fg-189 {
  color: #d7d7ff;
}
.asciinema-terminal .bg-189 {
  background-color: #d7d7ff;
}
.asciinema-terminal .fg-190 {
  color: #d7ff00;
}
.asciinema-terminal .bg-190 {
  background-color: #d7ff00;
}
.asciinema-terminal .fg-191 {
  color: #d7ff5f;
}
.asciinema-terminal .bg-191 {
  background-color: #d7ff5f;
}
.asciinema-terminal .fg-192 {
  color: #d7ff87;
}
.asciinema-terminal .bg-192 {
  background-color: #d7ff87;
}
.asciinema-terminal .fg-193 {
  color: #d7ffaf;
}
.asciinema-terminal .bg-193 {
  background-color: #d7ffaf;
}
.asciinema-terminal .fg-194 {
  color: #d7ffd7;
}
.asciinema-terminal .bg-194 {
  background-color: #d7ffd7;
}
.asciinema-terminal .fg-195 {
  color: #d7ffff;
}
.asciinema-terminal .bg-195 {
  background-color: #d7ffff;
}
.asciinema-terminal .fg-196 {
  color: #ff0000;
}
.asciinema-terminal .bg-196 {
  background-color: #ff0000;
}
.asciinema-terminal .fg-197 {
  color: #ff005f;
}
.asciinema-terminal .bg-197 {
  background-color: #ff005f;
}
.asciinema-terminal .fg-198 {
  color: #ff0087;
}
.asciinema-terminal .bg-198 {
  background-color: #ff0087;
}
.asciinema-terminal .fg-199 {
  color: #ff00af;
}
.asciinema-terminal .bg-199 {
  background-color: #ff00af;
}
.asciinema-terminal .fg-200 {
  color: #ff00d7;
}
.asciinema-terminal .bg-200 {
  background-color: #ff00d7;
}
.asciinema-terminal .fg-201 {
  color: #ff00ff;
}
.asciinema-terminal .bg-201 {
  background-color: #ff00ff;
}
.asciinema-terminal .fg-202 {
  color: #ff5f00;
}
.asciinema-terminal .bg-202 {
  background-color: #ff5f00;
}
.asciinema-terminal .fg-203 {
  color: #ff5f5f;
}
.asciinema-terminal .bg-203 {
  background-color: #ff5f5f;
}
.asciinema-terminal .fg-204 {
  color: #ff5f87;
}
.asciinema-terminal .bg-204 {
  background-color: #ff5f87;
}
.asciinema-terminal .fg-205 {
  color: #ff5faf;
}
.asciinema-terminal .bg-205 {
  background-color: #ff5faf;
}
.asciinema-terminal .fg-206 {
  color: #ff5fd7;
}
.asciinema-terminal .bg-206 {
  background-color: #ff5fd7;
}
.asciinema-terminal .fg-207 {
  color: #ff5fff;
}
.asciinema-terminal .bg-207 {
  background-color: #ff5fff;
}
.asciinema-terminal .fg-208 {
  color: #ff8700;
}
.asciinema-terminal .bg-208 {
  background-color: #ff8700;
}
.asciinema-terminal .fg-209 {
  color: #ff875f;
}
.asciinema-terminal .bg-209 {
  background-color: #ff875f;
}
.asciinema-terminal .fg-210 {
  color: #ff8787;
}
.asciinema-terminal .bg-210 {
  background-color: #ff8787;
}
.asciinema-terminal .fg-211 {
  color: #ff87af;
}
.asciinema-terminal .bg-211 {
  background-color: #ff87af;
}
.asciinema-terminal .fg-212 {
  color: #ff87d7;
}
.asciinema-terminal .bg-212 {
  background-color: #ff87d7;
}
.asciinema-terminal .fg-213 {
  color: #ff87ff;
}
.asciinema-terminal .bg-213 {
  background-color: #ff87ff;
}
.asciinema-terminal .fg-214 {
  color: #ffaf00;
}
.asciinema-terminal .bg-214 {
  background-color: #ffaf00;
}
.asciinema-terminal .fg-215 {
  color: #ffaf5f;
}
.asciinema-terminal .bg-215 {
  background-color: #ffaf5f;
}
.asciinema-terminal .fg-216 {
  color: #ffaf87;
}
.asciinema-terminal .bg-216 {
  background-color: #ffaf87;
}
.asciinema-terminal .fg-217 {
  color: #ffafaf;
}
.asciinema-terminal .bg-217 {
  background-color: #ffafaf;
}
.asciinema-terminal .fg-218 {
  color: #ffafd7;
}
.asciinema-terminal .bg-218 {
  background-color: #ffafd7;
}
.asciinema-terminal .fg-219 {
  color: #ffafff;
}
.asciinema-terminal .bg-219 {
  background-color: #ffafff;
}
.asciinema-terminal .fg-220 {
  color: #ffd700;
}
.asciinema-terminal .bg-220 {
  background-color: #ffd700;
}
.asciinema-terminal .fg-221 {
  color: #ffd75f;
}
.asciinema-terminal .bg-221 {
  background-color: #ffd75f;
}
.asciinema-terminal .fg-222 {
  color: #ffd787;
}
.asciinema-terminal .bg-222 {
  background-color: #ffd787;
}
.asciinema-terminal .fg-223 {
  color: #ffd7af;
}
.asciinema-terminal .bg-223 {
  background-color: #ffd7af;
}
.asciinema-terminal .fg-224 {
  color: #ffd7d7;
}
.asciinema-terminal .bg-224 {
  background-color: #ffd7d7;
}
.asciinema-terminal .fg-225 {
  color: #ffd7ff;
}
.asciinema-terminal .bg-225 {
  background-color: #ffd7ff;
}
.asciinema-terminal .fg-226 {
  color: #ffff00;
}
.asciinema-terminal .bg-226 {
  background-color: #ffff00;
}
.asciinema-terminal .fg-227 {
  color: #ffff5f;
}
.asciinema-terminal .bg-227 {
  background-color: #ffff5f;
}
.asciinema-terminal .fg-228 {
  color: #ffff87;
}
.asciinema-terminal .bg-228 {
  background-color: #ffff87;
}
.asciinema-terminal .fg-229 {
  color: #ffffaf;
}
.asciinema-terminal .bg-229 {
  background-color: #ffffaf;
}
.asciinema-terminal .fg-230 {
  color: #ffffd7;
}
.asciinema-terminal .bg-230 {
  background-color: #ffffd7;
}
.asciinema-terminal .fg-231 {
  color: #ffffff;
}
.asciinema-terminal .bg-231 {
  background-color: #ffffff;
}
.asciinema-terminal .fg-232 {
  color: #080808;
}
.asciinema-terminal .bg-232 {
  background-color: #080808;
}
.asciinema-terminal .fg-233 {
  color: #121212;
}
.asciinema-terminal .bg-233 {
  background-color: #121212;
}
.asciinema-terminal .fg-234 {
  color: #1c1c1c;
}
.asciinema-terminal .bg-234 {
  background-color: #1c1c1c;
}
.asciinema-terminal .fg-235 {
  color: #262626;
}
.asciinema-terminal .bg-235 {
  background-color: #262626;
}
.asciinema-terminal .fg-236 {
  color: #303030;
}
.asciinema-terminal .bg-236 {
  background-color: #303030;
}
.asciinema-terminal .fg-237 {
  color: #3a3a3a;
}
.asciinema-terminal .bg-237 {
  background-color: #3a3a3a;
}
.asciinema-terminal .fg-238 {
  color: #444444;
}
.asciinema-terminal .bg-238 {
  background-color: #444444;
}
.asciinema-terminal .fg-239 {
  color: #4e4e4e;
}
.asciinema-terminal .bg-239 {
  background-color: #4e4e4e;
}
.asciinema-terminal .fg-240 {
  color: #585858;
}
.asciinema-terminal .bg-240 {
  background-color: #585858;
}
.asciinema-terminal .fg-241 {
  color: #626262;
}
.asciinema-terminal .bg-241 {
  background-color: #626262;
}
.asciinema-terminal .fg-242 {
  color: #6c6c6c;
}
.asciinema-terminal .bg-242 {
  background-color: #6c6c6c;
}
.asciinema-terminal .fg-243 {
  color: #767676;
}
.asciinema-terminal .bg-243 {
  background-color: #767676;
}
.asciinema-terminal .fg-244 {
  color: #808080;
}
.asciinema-terminal .bg-244 {
  background-color: #808080;
}
.asciinema-terminal .fg-245 {
  color: #8a8a8a;
}
.asciinema-terminal .bg-245 {
  background-color: #8a8a8a;
}
.asciinema-terminal .fg-246 {
  color: #949494;
}
.asciinema-terminal .bg-246 {
  background-color: #949494;
}
.asciinema-terminal .fg-247 {
  color: #9e9e9e;
}
.asciinema-terminal .bg-247 {
  background-color: #9e9e9e;
}
.asciinema-terminal .fg-248 {
  color: #a8a8a8;
}
.asciinema-terminal .bg-248 {
  background-color: #a8a8a8;
}
.asciinema-terminal .fg-249 {
  color: #b2b2b2;
}
.asciinema-terminal .bg-249 {
  background-color: #b2b2b2;
}
.asciinema-terminal .fg-250 {
  color: #bcbcbc;
}
.asciinema-terminal .bg-250 {
  background-color: #bcbcbc;
}
.asciinema-terminal .fg-251 {
  color: #c6c6c6;
}
.asciinema-terminal .bg-251 {
  background-color: #c6c6c6;
}
.asciinema-terminal .fg-252 {
  color: #d0d0d0;
}
.asciinema-terminal .bg-252 {
  background-color: #d0d0d0;
}
.asciinema-terminal .fg-253 {
  color: #dadada;
}
.asciinema-terminal .bg-253 {
  background-color: #dadada;
}
.asciinema-terminal .fg-254 {
  color: #e4e4e4;
}
.asciinema-terminal .bg-254 {
  background-color: #e4e4e4;
}
.asciinema-terminal .fg-255 {
  color: #eeeeee;
}
.asciinema-terminal .bg-255 {
  background-color: #eeeeee;
}
.asciinema-theme-asciinema .asciinema-terminal {
  color: #cccccc;
  background-color: #121314;
  border-color: #121314;
}
.asciinema-theme-asciinema .fg-bg {
  color: #121314;
}
.asciinema-theme-asciinema .bg-fg {
  background-color: #cccccc;
}
.asciinema-theme-asciinema .fg-0 {
  color: #000000;
}
.asciinema-theme-asciinema .bg-0 {
  background-color: #000000;
}
.asciinema-theme-asciinema .fg-1 {
  color: #dd3c69;
}
.asciinema-theme-asciinema .bg-1 {
  background-color: #dd3c69;
}
.asciinema-theme-asciinema .fg-2 {
  color: #4ebf22;
}
.asciinema-theme-asciinema .bg-2 {
  background-color: #4ebf22;
}
.asciinema-theme-asciinema .fg-3 {
  color: #ddaf3c;
}
.asciinema-theme-asciinema .bg-3 {
  background-color: #ddaf3c;
}
.asciinema-theme-asciinema .fg-4 {
  color: #26b0d7;
}
.asciinema-theme-asciinema .bg-4 {
  background-color: #26b0d7;
}
.asciinema-theme-asciinema .fg-5 {
  color: #b954e1;
}
.asciinema-theme-asciinema .bg-5 {
  background-color: #b954e1;
}
.asciinema-theme-asciinema .fg-6 {
  color: #54e1b9;
}
.asciinema-theme-asciinema .bg-6 {
  background-color: #54e1b9;
}
.asciinema-theme-asciinema .fg-7 {
  color: #d9d9d9;
}
.asciinema-theme-asciinema .bg-7 {
  background-color: #d9d9d9;
}
.asciinema-theme-asciinema .fg-8 {
  color: #4d4d4d;
}
.asciinema-theme-asciinema .bg-8 {
  background-color: #4d4d4d;
}
.asciinema-theme-asciinema .fg-9 {
  color: #dd3c69;
}
.asciinema-theme-asciinema .bg-9 {
  background-color: #dd3c69;
}
.asciinema-theme-asciinema .fg-10 {
  color: #4ebf22;
}
.asciinema-theme-asciinema .bg-10 {
  background-color: #4ebf22;
}
.asciinema-theme-asciinema .fg-11 {
  color: #ddaf3c;
}
.asciinema-theme-asciinema .bg-11 {
  background-color: #ddaf3c;
}
.asciinema-theme-asciinema .fg-12 {
  color: #26b0d7;
}
.asciinema-theme-asciinema .bg-12 {
  background-color: #26b0d7;
}
.asciinema-theme-asciinema .fg-13 {
  color: #b954e1;
}
.asciinema-theme-asciinema .bg-13 {
  background-color: #b954e1;
}
.asciinema-theme-asciinema .fg-14 {
  color: #54e1b9;
}
.asciinema-theme-asciinema .bg-14 {
  background-color: #54e1b9;
}
.asciinema-theme-asciinema .fg-15 {
  color: #ffffff;
}
.asciinema-theme-asciinema .bg-15 {
  background-color: #ffffff;
}
.asciinema-theme-asciinema .fg-8,
.asciinema-theme-asciinema .fg-9,
.asciinema-theme-asciinema .fg-10,
.asciinema-theme-asciinema .fg-11,
.asciinema-theme-asciinema .fg-12,
.asciinema-theme-asciinema .fg-13,
.asciinema-theme-asciinema .fg-14,
.asciinema-theme-asciinema .fg-15 {
  font-weight: bold;
}
.asciinema-theme-tango .asciinema-terminal {
  color: #cccccc;
  background-color: #121314;
  border-color: #121314;
}
.asciinema-theme-tango .fg-bg {
  color: #121314;
}
.asciinema-theme-tango .bg-fg {
  background-color: #cccccc;
}
.asciinema-theme-tango .fg-0 {
  color: #000000;
}
.asciinema-theme-tango .bg-0 {
  background-color: #000000;
}
.asciinema-theme-tango .fg-1 {
  color: #cc0000;
}
.asciinema-theme-tango .bg-1 {
  background-color: #cc0000;
}
.asciinema-theme-tango .fg-2 {
  color: #4e9a06;
}
.asciinema-theme-tango .bg-2 {
  background-color: #4e9a06;
}
.asciinema-theme-tango .fg-3 {
  color: #c4a000;
}
.asciinema-theme-tango .bg-3 {
  background-color: #c4a000;
}
.asciinema-theme-tango .fg-4 {
  color: #3465a4;
}
.asciinema-theme-tango .bg-4 {
  background-color: #3465a4;
}
.asciinema-theme-tango .fg-5 {
  color: #75507b;
}
.asciinema-theme-tango .bg-5 {
  background-color: #75507b;
}
.asciinema-theme-tango .fg-6 {
  color: #06989a;
}
.asciinema-theme-tango .bg-6 {
  background-color: #06989a;
}
.asciinema-theme-tango .fg-7 {
  color: #d3d7cf;
}
.asciinema-theme-tango .bg-7 {
  background-color: #d3d7cf;
}
.asciinema-theme-tango .fg-8 {
  color: #555753;
}
.asciinema-theme-tango .bg-8 {
  background-color: #555753;
}
.asciinema-theme-tango .fg-9 {
  color: #ef2929;
}
.asciinema-theme-tango .bg-9 {
  background-color: #ef2929;
}
.asciinema-theme-tango .fg-10 {
  color: #8ae234;
}
.asciinema-theme-tango .bg-10 {
  background-color: #8ae234;
}
.asciinema-theme-tango .fg-11 {
  color: #fce94f;
}
.asciinema-theme-tango .bg-11 {
  background-color: #fce94f;
}
.asciinema-theme-tango .fg-12 {
  color: #729fcf;
}
.asciinema-theme-tango .bg-12 {
  background-color: #729fcf;
}
.asciinema-theme-tango .fg-13 {
  color: #ad7fa8;
}
.asciinema-theme-tango .bg-13 {
  background-color: #ad7fa8;
}
.asciinema-theme-tango .fg-14 {
  color: #34e2e2;
}
.asciinema-theme-tango .bg-14 {
  background-color: #34e2e2;
}
.asciinema-theme-tango .fg-15 {
  color: #eeeeec;
}
.asciinema-theme-tango .bg-15 {
  background-color: #eeeeec;
}
.asciinema-theme-tango .fg-8,
.asciinema-theme-tango .fg-9,
.asciinema-theme-tango .fg-10,
.asciinema-theme-tango .fg-11,
.asciinema-theme-tango .fg-12,
.asciinema-theme-tango .fg-13,
.asciinema-theme-tango .fg-14,
.asciinema-theme-tango .fg-15 {
  font-weight: bold;
}
.asciinema-theme-solarized-dark .asciinema-terminal {
  color: #839496;
  background-color: #002b36;
  border-color: #002b36;
}
.asciinema-theme-solarized-dark .fg-bg {
  color: #002b36;
}
.asciinema-theme-solarized-dark .bg-fg {
  background-color: #839496;
}
.asciinema-theme-solarized-dark .fg-0 {
  color: #073642;
}
.asciinema-theme-solarized-dark .bg-0 {
  background-color: #073642;
}
.asciinema-theme-solarized-dark .fg-1 {
  color: #dc322f;
}
.asciinema-theme-solarized-dark .bg-1 {
  background-color: #dc322f;
}
.asciinema-theme-solarized-dark .fg-2 {
  color: #859900;
}
.asciinema-theme-solarized-dark .bg-2 {
  background-color: #859900;
}
.asciinema-theme-solarized-dark .fg-3 {
  color: #b58900;
}
.asciinema-theme-solarized-dark .bg-3 {
  background-color: #b58900;
}
.asciinema-theme-solarized-dark .fg-4 {
  color: #268bd2;
}
.asciinema-theme-solarized-dark .bg-4 {
  background-color: #268bd2;
}
.asciinema-theme-solarized-dark .fg-5 {
  color: #d33682;
}
.asciinema-theme-solarized-dark .bg-5 {
  background-color: #d33682;
}
.asciinema-theme-solarized-dark .fg-6 {
  color: #2aa198;
}
.asciinema-theme-solarized-dark .bg-6 {
  background-color: #2aa198;
}
.asciinema-theme-solarized-dark .fg-7 {
  color: #eee8d5;
}
.asciinema-theme-solarized-dark .bg-7 {
  background-color: #eee8d5;
}
.asciinema-theme-solarized-dark .fg-8 {
  color: #002b36;
}
.asciinema-theme-solarized-dark .bg-8 {
  background-color: #002b36;
}
.asciinema-theme-solarized-dark .fg-9 {
  color: #cb4b16;
}
.asciinema-theme-solarized-dark .bg-9 {
  background-color: #cb4b16;
}
.asciinema-theme-solarized-dark .fg-10 {
  color: #586e75;
}
.asciinema-theme-solarized-dark .bg-10 {
  background-color: #586e75;
}
.asciinema-theme-solarized-dark .fg-11 {
  color: #657b83;
}
.asciinema-theme-solarized-dark .bg-11 {
  background-color: #657b83;
}
.asciinema-theme-solarized-dark .fg-12 {
  color: #839496;
}
.asciinema-theme-solarized-dark .bg-12 {
  background-color: #839496;
}
.asciinema-theme-solarized-dark .fg-13 {
  color: #6c71c4;
}
.asciinema-theme-solarized-dark .bg-13 {
  background-color: #6c71c4;
}
.asciinema-theme-solarized-dark .fg-14 {
  color: #93a1a1;
}
.asciinema-theme-solarized-dark .bg-14 {
  background-color: #93a1a1;
}
.asciinema-theme-solarized-dark .fg-15 {
  color: #fdf6e3;
}
.asciinema-theme-solarized-dark .bg-15 {
  background-color: #fdf6e3;
}
.asciinema-theme-solarized-light .asciinema-terminal {
  color: #657b83;
  background-color: #fdf6e3;
  border-color: #fdf6e3;
}
.asciinema-theme-solarized-light .fg-bg {
  color: #fdf6e3;
}
.asciinema-theme-solarized-light .bg-fg {
  background-color: #657b83;
}
.asciinema-theme-solarized-light .fg-0 {
  color: #073642;
}
.asciinema-theme-solarized-light .bg-0 {
  background-color: #073642;
}
.asciinema-theme-solarized-light .fg-1 {
  color: #dc322f;
}
.asciinema-theme-solarized-light .bg-1 {
  background-color: #dc322f;
}
.asciinema-theme-solarized-light .fg-2 {
  color: #859900;
}
.asciinema-theme-solarized-light .bg-2 {
  background-color: #859900;
}
.asciinema-theme-solarized-light .fg-3 {
  color: #b58900;
}
.asciinema-theme-solarized-light .bg-3 {
  background-color: #b58900;
}
.asciinema-theme-solarized-light .fg-4 {
  color: #268bd2;
}
.asciinema-theme-solarized-light .bg-4 {
  background-color: #268bd2;
}
.asciinema-theme-solarized-light .fg-5 {
  color: #d33682;
}
.asciinema-theme-solarized-light .bg-5 {
  background-color: #d33682;
}
.asciinema-theme-solarized-light .fg-6 {
  color: #2aa198;
}
.asciinema-theme-solarized-light .bg-6 {
  background-color: #2aa198;
}
.asciinema-theme-solarized-light .fg-7 {
  color: #eee8d5;
}
.asciinema-theme-solarized-light .bg-7 {
  background-color: #eee8d5;
}
.asciinema-theme-solarized-light .fg-8 {
  color: #002b36;
}
.asciinema-theme-solarized-light .bg-8 {
  background-color: #002b36;
}
.asciinema-theme-solarized-light .fg-9 {
  color: #cb4b16;
}
.asciinema-theme-solarized-light .bg-9 {
  background-color: #cb4b16;
}
.asciinema-theme-solarized-light .fg-10 {
  color: #586e75;
}
.asciinema-theme-solarized-light .bg-10 {
  background-color: #586e75;
}
.asciinema-theme-solarized-light .fg-11 {
  color: #657c83;
}
.asciinema-theme-solarized-light .bg-11 {
  background-color: #657c83;
}
.asciinema-theme-solarized-light .fg-12 {
  color: #839496;
}
.asciinema-theme-solarized-light .bg-12 {
  background-color: #839496;
}
.asciinema-theme-solarized-light .fg-13 {
  color: #6c71c4;
}
.asciinema-theme-solarized-light .bg-13 {
  background-color: #6c71c4;
}
.asciinema-theme-solarized-light .fg-14 {
  color: #93a1a1;
}
.asciinema-theme-solarized-light .bg-14 {
  background-color: #93a1a1;
}
.asciinema-theme-solarized-light .fg-15 {
  color: #fdf6e3;
}
.asciinema-theme-solarized-light .bg-15 {
  background-color: #fdf6e3;
}
.asciinema-theme-seti .asciinema-terminal {
  color: #cacecd;
  background-color: #111213;
  border-color: #111213;
}
.asciinema-theme-seti .fg-bg {
  color: #111213;
}
.asciinema-theme-seti .bg-fg {
  background-color: #cacecd;
}
.asciinema-theme-seti .fg-0 {
  color: #323232;
}
.asciinema-theme-seti .bg-0 {
  background-color: #323232;
}
.asciinema-theme-seti .fg-1 {
  color: #c22832;
}
.asciinema-theme-seti .bg-1 {
  background-color: #c22832;
}
.asciinema-theme-seti .fg-2 {
  color: #8ec43d;
}
.asciinema-theme-seti .bg-2 {
  background-color: #8ec43d;
}
.asciinema-theme-seti .fg-3 {
  color: #e0c64f;
}
.asciinema-theme-seti .bg-3 {
  background-color: #e0c64f;
}
.asciinema-theme-seti .fg-4 {
  color: #43a5d5;
}
.asciinema-theme-seti .bg-4 {
  background-color: #43a5d5;
}
.asciinema-theme-seti .fg-5 {
  color: #8b57b5;
}
.asciinema-theme-seti .bg-5 {
  background-color: #8b57b5;
}
.asciinema-theme-seti .fg-6 {
  color: #8ec43d;
}
.asciinema-theme-seti .bg-6 {
  background-color: #8ec43d;
}
.asciinema-theme-seti .fg-7 {
  color: #eeeeee;
}
.asciinema-theme-seti .bg-7 {
  background-color: #eeeeee;
}
.asciinema-theme-seti .fg-8 {
  color: #323232;
}
.asciinema-theme-seti .bg-8 {
  background-color: #323232;
}
.asciinema-theme-seti .fg-9 {
  color: #c22832;
}
.asciinema-theme-seti .bg-9 {
  background-color: #c22832;
}
.asciinema-theme-seti .fg-10 {
  color: #8ec43d;
}
.asciinema-theme-seti .bg-10 {
  background-color: #8ec43d;
}
.asciinema-theme-seti .fg-11 {
  color: #e0c64f;
}
.asciinema-theme-seti .bg-11 {
  background-color: #e0c64f;
}
.asciinema-theme-seti .fg-12 {
  color: #43a5d5;
}
.asciinema-theme-seti .bg-12 {
  background-color: #43a5d5;
}
.asciinema-theme-seti .fg-13 {
  color: #8b57b5;
}
.asciinema-theme-seti .bg-13 {
  background-color: #8b57b5;
}
.asciinema-theme-seti .fg-14 {
  color: #8ec43d;
}
.asciinema-theme-seti .bg-14 {
  background-color: #8ec43d;
}
.asciinema-theme-seti .fg-15 {
  color: #ffffff;
}
.asciinema-theme-seti .bg-15 {
  background-color: #ffffff;
}
.asciinema-theme-seti .fg-8,
.asciinema-theme-seti .fg-9,
.asciinema-theme-seti .fg-10,
.asciinema-theme-seti .fg-11,
.asciinema-theme-seti .fg-12,
.asciinema-theme-seti .fg-13,
.asciinema-theme-seti .fg-14,
.asciinema-theme-seti .fg-15 {
  font-weight: bold;
}
/* Based on Monokai from base16 collection - https://github.com/chriskempson/base16 */
.asciinema-theme-monokai .asciinema-terminal {
  color: #f8f8f2;
  background-color: #272822;
  border-color: #272822;
}
.asciinema-theme-monokai .fg-bg {
  color: #272822;
}
.asciinema-theme-monokai .bg-fg {
  background-color: #f8f8f2;
}
.asciinema-theme-monokai .fg-0 {
  color: #272822;
}
.asciinema-theme-monokai .bg-0 {
  background-color: #272822;
}
.asciinema-theme-monokai .fg-1 {
  color: #f92672;
}
.asciinema-theme-monokai .bg-1 {
  background-color: #f92672;
}
.asciinema-theme-monokai .fg-2 {
  color: #a6e22e;
}
.asciinema-theme-monokai .bg-2 {
  background-color: #a6e22e;
}
.asciinema-theme-monokai .fg-3 {
  color: #f4bf75;
}
.asciinema-theme-monokai .bg-3 {
  background-color: #f4bf75;
}
.asciinema-theme-monokai .fg-4 {
  color: #66d9ef;
}
.asciinema-theme-monokai .bg-4 {
  background-color: #66d9ef;
}
.asciinema-theme-monokai .fg-5 {
  color: #ae81ff;
}
.asciinema-theme-monokai .bg-5 {
  background-color: #ae81ff;
}
.asciinema-theme-monokai .fg-6 {
  color: #a1efe4;
}
.asciinema-theme-monokai .bg-6 {
  background-color: #a1efe4;
}
.asciinema-theme-monokai .fg-7 {
  color: #f8f8f2;
}
.asciinema-theme-monokai .bg-7 {
  background-color: #f8f8f2;
}
.asciinema-theme-monokai .fg-8 {
  color: #75715e;
}
.asciinema-theme-monokai .bg-8 {
  background-color: #75715e;
}
.asciinema-theme-monokai .fg-9 {
  color: #f92672;
}
.asciinema-theme-monokai .bg-9 {
  background-color: #f92672;
}
.asciinema-theme-monokai .fg-10 {
  color: #a6e22e;
}
.asciinema-theme-monokai .bg-10 {
  background-color: #a6e22e;
}
.asciinema-theme-monokai .fg-11 {
  color: #f4bf75;
}
.asciinema-theme-monokai .bg-11 {
  background-color: #f4bf75;
}
.asciinema-theme-monokai .fg-12 {
  color: #66d9ef;
}
.asciinema-theme-monokai .bg-12 {
  background-color: #66d9ef;
}
.asciinema-theme-monokai .fg-13 {
  color: #ae81ff;
}
.asciinema-theme-monokai .bg-13 {
  background-color: #ae81ff;
}
.asciinema-theme-monokai .fg-14 {
  color: #a1efe4;
}
.asciinema-theme-monokai .bg-14 {
  background-color: #a1efe4;
}
.asciinema-theme-monokai .fg-15 {
  color: #f9f8f5;
}
.asciinema-theme-monokai .bg-15 {
  background-color: #f9f8f5;
}
.asciinema-theme-monokai .fg-8,
.asciinema-theme-monokai .fg-9,
.asciinema-theme-monokai .fg-10,
.asciinema-theme-monokai .fg-11,
.asciinema-theme-monokai .fg-12,
.asciinema-theme-monokai .fg-13,
.asciinema-theme-monokai .fg-14,
.asciinema-theme-monokai .fg-15 {
  font-weight: bold;
}
`
