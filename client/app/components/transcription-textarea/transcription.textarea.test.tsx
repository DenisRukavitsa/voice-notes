import "@testing-library/jest-dom";
import { fireEvent, render, screen } from "@testing-library/react";
import TranscriptionTextarea from "./transcription.textarea";

describe("AudioTranscriber", () => {
  it("shows the passed transcription as the textarea value", () => {
    render(
      <TranscriptionTextarea transcription="test" onChange={() => null} />
    );

    const textarea = screen.getByRole("textbox");
    expect(textarea).toBeInTheDocument();
    expect(textarea).toHaveTextContent("test");
  });

  it("trigers onChange when the textarea value changes", () => {
    const handleChange = jest.fn();
    render(<TranscriptionTextarea transcription="" onChange={handleChange} />);

    const textarea = screen.getByRole("textbox") as HTMLTextAreaElement;
    fireEvent.change(textarea, { target: { value: "new value" } });
    expect(handleChange.mock.calls).toHaveLength(1);
    expect(handleChange.mock.calls[0][0]).toBe("new value");
  });
});
