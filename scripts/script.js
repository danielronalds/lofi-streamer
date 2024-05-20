const sendIframeCommand = (command) => {
  let iframe = document.getElementById("iframe");
  iframe.contentWindow.postMessage(
    JSON.stringify({ event: "command", func: command }),
    "*",
  );
};

const pauseVideo = () => {
  sendIframeCommand("stopVideo");
}

const playVideo = () => {
  sendIframeCommand("playVideo");
}
