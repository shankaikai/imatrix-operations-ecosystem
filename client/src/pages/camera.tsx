import type { NextPage } from "next";
import Stream from "../components/Camera/Stream";

const Camera: NextPage = () => {
  return (
    <div>
      <h1>Camera</h1>
      <Stream src="https://devstreaming-cdn.apple.com/videos/streaming/examples/img_bipbop_adv_example_fmp4/master.m3u8" id="1" />
    </div>
  );
};

export default Camera;
