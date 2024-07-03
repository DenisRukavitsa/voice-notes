import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";
import CopyNoteButton from "./copy.note.button";

describe("CopyNoteButton", () => {
  it("renders button in disabled state if note text is not provided", () => {
    render(<CopyNoteButton noteText="" />);
    const button = screen.getByRole("button");
    expect(button).toBeInTheDocument();
    expect(button).toHaveTextContent("Copy note");
    expect(button).toHaveAttribute("disabled");
  });
});
