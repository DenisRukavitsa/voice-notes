import "@testing-library/jest-dom";
import { render, screen } from "@testing-library/react";
import CopyNoteButton from "./copy.note.button";

describe("CopyNoteButton", () => {
  it("renders", () => {
    render(<CopyNoteButton />);
    const button = screen.getByRole("button");
    expect(button).toBeInTheDocument();
    expect(button).toHaveTextContent("Copy note");
  });
});
