import Hls from "hls.js";
import { FunctionComponent, useEffect, useState } from "react";

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
        hls.on(Hls.Events.ERROR, () => {
          console.log("error");
          setError(true);
        });
        //TODO: Add in plazceholder
        hls.on(Hls.Events.MANIFEST_PARSED, () => video.play());
      });
    }
  }, []);
  return (
    <>
      {error ? (
        <div
          style={{
            width: "100%",
            height: "100px",
          }}
        >
          Error
        </div>
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
