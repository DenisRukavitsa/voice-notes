import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";
import { act } from "react";
import RecordNoteButton from "./record.note.button";

describe("TakeNoteButton", () => {
  it("renders a button in initial state", () => {
    const handleRecordingStart = jest.fn();
    const handleRecordingStop = jest.fn();
    render(
      <RecordNoteButton
        onRecordingStart={handleRecordingStart}
        onRecordingStop={handleRecordingStop}
      />
    );

    const button = screen.getByRole("button");
    expect(button).toBeInTheDocument();
    expect(button).toHaveTextContent("Record a note");
    expect(handleRecordingStart.mock.calls).toHaveLength(0);
    expect(handleRecordingStop.mock.calls).toHaveLength(0);
  });

  it("changes button state to recording on click", async () => {
    const handleRecordingStart = jest.fn();
    const handleRecordingStop = jest.fn();
    render(
      <RecordNoteButton
        onRecordingStart={handleRecordingStart}
        onRecordingStop={handleRecordingStop}
      />
    );

    let button = screen.getByRole("button");
    act(() => button.click());
    expect(button).toHaveTextContent("Stop recording");
    expect(handleRecordingStart.mock.calls).toHaveLength(1);
    expect(handleRecordingStop.mock.calls).toHaveLength(0);
  });

  it("changes button state back to waiting after clicking second time", () => {
    const handleRecordingStart = jest.fn();
    const handleRecordingStop = jest.fn();
    render(
      <RecordNoteButton
        onRecordingStart={handleRecordingStart}
        onRecordingStop={handleRecordingStop}
      />
    );

    const button = screen.getByRole("button");
    act(() => button.click());
    act(() => button.click());
    expect(button).toHaveTextContent("Record a note");
    expect(handleRecordingStart.mock.calls).toHaveLength(1);
    expect(handleRecordingStop.mock.calls).toHaveLength(1);
  });
});
