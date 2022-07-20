import Hls from "hls.js";
import { FunctionComponent, useEffect, useState } from "react";
import { Text } from "@mantine/core";

interface StreamProps {
  src: string;
  id: any;
}

const Stream: FunctionComponent<StreamProps> = (props: StreamProps) => {
  const [error, setError] = useState(false);

  useEffect(() => {
    if (Hls.isSupported()) {
      let video = document.getElementById(
        props.id.toString()
      ) as HTMLMediaElement;
      let hls = new Hls();
      // bind them together
      hls.attachMedia(video);
      hls.on(Hls.Events.MEDIA_ATTACHED, () => {
        hls.loadSource(props.src);
        // hls.on(Hls.Events.ERROR, () => {
        //   console.log("hls error");
        //   setError(true);
        // });
        //TODO: Add in plazceholder
        hls.on(Hls.Events.MANIFEST_PARSED, () => {
          // setError(false);
          video.play();
        });
      });
    }
  }, []);
  return (
    <>
      {error ? (
        <Text
          style={{
            width: "100%",
            height: "100px",
          }}
        >
          Error
        </Text>
      ) : (
        <video
          id={props.id.toString()}
          style={{
            width: "100%",
          }}
        />
      )}
    </>
  );
};

export default Stream;
