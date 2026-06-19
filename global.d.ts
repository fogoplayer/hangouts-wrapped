type HTMLEvent<El = HTMLElement> = Event & {
  target: El;
  currentTarget: El;
};
type HTMLInputEvent = HTMLEvent<HTMLInputElement>;
type SlotChangeEvent = HTMLEvent<HTMLSlotElement> | { target: HTMLSlotElement };
type HTMLMouseEvent<El = HTMLElement> = MouseEvent & HTMLEvent<El>;
type UpdatedDiff = Map<PropertyKey, unknown>;
