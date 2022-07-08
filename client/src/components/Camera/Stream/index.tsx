
import Hls from "hls.js";
import { FunctionComponent, useEffect } from "react";

interface StreamProps {
    src:string
    id:any
}
 
const Stream: FunctionComponent<StreamProps> = (props:StreamProps) => {
    useEffect(()=>{
        if (Hls.isSupported()) {
            let video = document.getElementById(props.id.toString()) as HTMLMediaElement;
            let hls = new Hls();
            // bind them together
            hls.attachMedia(video);
            hls.loadSource(props.src)
          }
    },[])
    return (
        <>
        <video id={props.id.toString()} />
        </>
    );
}
 
export default Stream;