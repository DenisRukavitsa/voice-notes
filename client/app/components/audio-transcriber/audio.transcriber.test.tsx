import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";
import AudioTranscriber from "./audio.transcriber";
import { act } from "react";

interface MediaDevicesProp {
  value: {
    getUserMedia?: () => void;
  };
}

Object.defineProperty(window, "MediaRecorder", {
  value: jest.fn().mockImplementation(() => ({
    start: jest.fn(),
    onstop: jest.fn(),
    ondataavailable: jest.fn(),
    stop: function () {
      this.onstop();
    },
    stream: {
      getTracks: () => [
        {
          stop: jest.fn(),
        },
      ],
    },
  })),
});

global.fetch = jest.fn(() =>
  Promise.resolve({
    ok: true,
    json: () => Promise.resolve({ transcription: "test" }),
  })
) as jest.Mock;

describe("AudioTranscriber", () => {
  const mediaDevices: MediaDevicesProp = {
    value: { getUserMedia: () => null },
  };
  Object.defineProperty(global.navigator, "mediaDevices", mediaDevices);

  beforeEach(() => {
    mediaDevices.value.getUserMedia = () => null;
  });

  it("disables the Record note button and returns an error if navigator.mediaDevices.getUserMedia is not defined", () => {
    mediaDevices.value.getUserMedia = undefined;

    const onError = jest.fn();
    render(
      <AudioTranscriber
        onError={onError}
        onTranscription={() => null}
        clearError={() => null}
      />
    );

    const button = screen.getByRole("button");
    expect(button).toHaveAttribute("disabled");
    expect(onError.mock.calls).toHaveLength(1);
    expect(onError.mock.calls[0][0]).toBe(
      "Audio input is not available on this device"
    );
  });

  it("clears the error when recording is started ", async () => {
    const clearError = jest.fn();
    render(
      <AudioTranscriber
        onError={() => null}
        onTranscription={() => null}
        clearError={clearError}
      />
    );

    const button = screen.getByRole("button");
    await act(async () => await button.click());
    expect(clearError.mock.calls).toHaveLength(1);
  });

  it("returns a trancription when recording is stopped", async () => {
    const onTranscription = jest.fn();
    render(
      <AudioTranscriber
        onError={() => null}
        onTranscription={onTranscription}
        clearError={() => null}
      />
    );

    const button = screen.getByRole("button");
    await act(async () => await button.click());
    await act(async () => await button.click());
    expect(onTranscription.mock.calls).toHaveLength(1);
    expect(onTranscription.mock.calls[0][0]).toBe("test");
  });
});
