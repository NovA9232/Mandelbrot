#include "raylib.h"

int mandelbrotPixel(double i, double j, int W, int H, int maxIteration, double zoom) {
	double zx = ((i/double(W))*2.5)/zoom;
	double zy = ((j/double(H))*1.5)/zoom;
	double x = 0.0;
	double y = 0.0;
	int iteration = 0;

	while ((x*x + y*y < 4) && (iteration < maxIteration)) {
		double xtemp = x*x - y*y + zx;
		y = (x * y * 2) + zy;
		x = xtemp;
		iteration++;
	}

	return iteration;
}


void fullMandelbrot(double offX, double offY, int W, int H, double zoom, int maxIteration) {
	for (double i = 0; i < W; i++) {
		for (double j = 0; j < H; j++) {
			int iteration = mandelbrotPixel(i + offX, j + offY, W, H, maxIteration, zoom);
			unsigned char col = (double(iteration)/double(maxIteration)) * 255;
			DrawPixel(i, j, Color{col, col, col, 255});
		}
	}
}


int main(int argc, char* argv[]) {
	int screenWidth = 1000;
	int screenHeight = 600;

	InitWindow(screenWidth, screenHeight, "Mandelbrot");
	SetTargetFPS(144);

	double offX = -800;
	double offY = -300;

	bool rendered = false;

	while (!WindowShouldClose()) {

		if (IsMouseButtonDown(0)) {
			offX += (double(GetMouseX()) - double(screenWidth/2));
			offY += (double(GetMouseY()) - double(screenHeight/2));
			rendered = false;
		}

		BeginDrawing();
		if (!rendered) {
			ClearBackground(BLACK);
			fullMandelbrot(offX, offY, screenWidth, screenHeight, 1, 100);
			rendered = true;
		}

		EndDrawing();
	}
	CloseWindow();        // Close window and OpenGL context

	return 0;
}
