#!/usr/bin/env python3

from pathlib import Path
import argparse

try:
    import cairosvg
except ImportError:
    cairosvg = None


def generate_logo_svg(
    canvas_size: int = 1024,
    radius: float = 128,
    arm_length: float | None = None,
    stroke_width: float = 18,
    mark_color: str = "#C68A2C",
    background_color: str = "#050505",
) -> str:
    """
    Generate the 'Autonomous unit, infinite scale' logo as SVG.

    Geometry:
      R = circle radius
      L = arm length
      P = 2L/3 -> cross-cut position from circle edge
      W = L/3  -> cross-cut width
      E = L/3  -> remaining outer stub after cross-cut
    """

    if arm_length is None:
        arm_length = radius

    R = radius
    L = arm_length
    P = 2 * L / 3
    W = L / 3

    cx = canvas_size / 2
    cy = canvas_size / 2

    half_cut = W / 2

    svg = f'''<svg
  xmlns="http://www.w3.org/2000/svg"
  width="{canvas_size}"
  height="{canvas_size}"
  viewBox="0 0 {canvas_size} {canvas_size}"
  role="img"
  aria-label="Autonomous unit infinite scale logo"
>
  <rect width="{canvas_size}" height="{canvas_size}" fill="{background_color}"/>

  <!--
    Geometry spec

    Center: ({cx}, {cy})

    R  = {R}
    L  = {L}
    P  = 2L/3 = {P}
    E  = L/3  = {L / 3}
    W  = L/3  = {W}

    Circle radius        = R
    Arm length           = L
    Cross-cut position   = 2L/3 from circle edge
    Remaining outer stub = L/3
    Cross-cut width      = L/3

    Stroke width = {stroke_width}
    Color        = {mark_color}
  -->

  <defs>
    <style>
      .mark {{
        stroke: {mark_color};
        stroke-width: {stroke_width};
        stroke-linecap: square;
        stroke-linejoin: miter;
        fill: none;
        vector-effect: non-scaling-stroke;
        shape-rendering: geometricPrecision;
      }}
    </style>
  </defs>

  <g class="mark" transform="translate({cx} {cy})">
    <!-- central autonomous unit -->
    <circle cx="0" cy="0" r="{R}"/>

    <!-- top interface -->
    <line x1="0" y1="-{R}" x2="0" y2="-{R + L}"/>
    <line x1="-{half_cut}" y1="-{R + P}" x2="{half_cut}" y2="-{R + P}"/>

    <!-- right interface -->
    <line x1="{R}" y1="0" x2="{R + L}" y2="0"/>
    <line x1="{R + P}" y1="-{half_cut}" x2="{R + P}" y2="{half_cut}"/>

    <!-- bottom interface -->
    <line x1="0" y1="{R}" x2="0" y2="{R + L}"/>
    <line x1="-{half_cut}" y1="{R + P}" x2="{half_cut}" y2="{R + P}"/>

    <!-- left interface -->
    <line x1="-{R}" y1="0" x2="-{R + L}" y2="0"/>
    <line x1="-{R + P}" y1="-{half_cut}" x2="-{R + P}" y2="{half_cut}"/>
  </g>
</svg>
'''
    return svg


def main() -> None:
    parser = argparse.ArgumentParser(
        description="Render the autonomous unit logo as SVG and optionally PNG."
    )

    parser.add_argument("--out", default="autonomous_unit_logo", help="Output filename without extension")
    parser.add_argument("--size", type=int, default=1024, help="Canvas size in pixels")
    parser.add_argument("--radius", type=float, default=128, help="Circle radius R")
    parser.add_argument("--arm", type=float, default=None, help="Arm length L. Defaults to R")
    parser.add_argument("--stroke", type=float, default=18, help="Stroke width")
    parser.add_argument("--mark", default="#C68A2C", help="Logo color")
    parser.add_argument("--bg", default="#050505", help="Background color")
    parser.add_argument("--png", action="store_true", help="Also render PNG using cairosvg")

    args = parser.parse_args()

    svg = generate_logo_svg(
        canvas_size=args.size,
        radius=args.radius,
        arm_length=args.arm,
        stroke_width=args.stroke,
        mark_color=args.mark,
        background_color=args.bg,
    )

    svg_path = Path(f"{args.out}.svg")
    svg_path.write_text(svg, encoding="utf-8")

    print(f"SVG written to: {svg_path}")

    if args.png:
        if cairosvg is None:
            raise RuntimeError("PNG export requires cairosvg. Install with: pip install cairosvg")

        png_path = Path(f"{args.out}.png")
        cairosvg.svg2png(
            bytestring=svg.encode("utf-8"),
            write_to=str(png_path),
            output_width=args.size,
            output_height=args.size,
        )

        print(f"PNG written to: {png_path}")


if __name__ == "__main__":
    main()
