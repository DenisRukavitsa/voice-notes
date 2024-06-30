import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";
import AudioTranscriber from "./audio.transcriber";

describe("AudioTranscriber", () => {
  it("disbales the Record note button if navigator.mediaDevices is undefined", () => {
    render(
      <AudioTranscriber onTranscription={() => null} onError={() => null} />
    );
    const button = screen.getByRole("button");
    expect(button).toHaveAttribute("disabled");
  });
});
