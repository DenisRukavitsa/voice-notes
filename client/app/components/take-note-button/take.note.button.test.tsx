import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";
import TakeNoteButton from "./take.note.button";

describe("TakeNoteButton", () => {
  it("renders", () => {
    render(<TakeNoteButton />);
    const button = screen.getByRole("button");
    expect(button).toBeInTheDocument();
    expect(button).toHaveTextContent("Take a note");
  });
});
