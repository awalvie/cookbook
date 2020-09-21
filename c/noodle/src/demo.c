#include <stdbool.h>
#include <SDL2/SDL.h>
#include <stdio.h>

#define SCREEN_WIDTH 640
#define SCREEN_HEIGHT 480

SDL_Window *gWindow = NULL;
SDL_Renderer *renderer = NULL;
SDL_Surface *gSurface = NULL;
SDL_Surface *gCanvas = NULL;
bool gQuit = false;

int error(char *msg, const char *err)
{
	printf("Error %s: %s\n", msg, err);
}

bool init()
{
	bool success = true;
	if (SDL_Init(SDL_INIT_VIDEO) < 0) {
		error("init", SDL_GetError());
		success = false;
	} else {
		gWindow =
			SDL_CreateWindow("noodle", SDL_WINDOWPOS_UNDEFINED,
					 SDL_WINDOWPOS_UNDEFINED, SCREEN_WIDTH,
					 SCREEN_HEIGHT, SDL_WINDOW_SHOWN);

		if (gWindow == NULL) {
			error("window", SDL_GetError());
			success = false;
		} else {
			gSurface = SDL_GetWindowSurface(gWindow);
			if (gSurface == NULL) {
				error("surface", SDL_GetError());
				success = false;
			}
		}
	}

	return success;
}

bool loadMedia()
{
	bool success = true;

	gCanvas = SDL_CreateRGBSurface(0, SCREEN_WIDTH, SCREEN_HEIGHT, 32, 0, 0,
				       0, 0);
	SDL_FillRect(gCanvas, NULL,
		     SDL_MapRGB(gCanvas->format, 0xFF, 0xFF, 0xFF));

	if (gCanvas == NULL) {
		error("canvas", SDL_GetError());
		success = false;
	}

	return success;
}

bool quit()
{
	SDL_FreeSurface(gCanvas);
	gCanvas = NULL;
	SDL_DestroyWindow(gWindow);
	gWindow = NULL;
	SDL_Quit();
}

void end_loop()
{
	gQuit = true;
}

void handle_mouse(SDL_Event *e)
{
	int mouseX, mouseY;

	switch (e->type) {
	case SDL_MOUSEMOTION:
		SDL_GetMouseState(&mouseX, &mouseY);
		printf("X: %d\nY: %d\n", mouseX, mouseY);
		break;
	}
}

void handle_keypress(SDL_Event *event)
{
	switch (event->key.keysym.sym) {
	case SDLK_q:
		end_loop();
		break;
	}
}

int main()
{
	if (!init()) {
		fprintf(stderr, "ERROR: Could not initialize SDL\n");
	} else {
		SDL_BlitSurface(gCanvas, NULL, gSurface, NULL);
		SDL_UpdateWindowSurface(gWindow);
		/* main game loop */
		SDL_Event event;

		while (!gQuit) {
			while (SDL_PollEvent(&event) != 0) {
				if (event.type == SDL_MOUSEMOTION) {
					handle_mouse(&event);
				} else if (event.type == SDL_KEYDOWN) {
					handle_keypress(&event);
				}
			}
		}
	}

	quit();
	return 0;
}
