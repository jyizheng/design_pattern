#include <iostream>

// Abstract Class
class Game {
public:
    // Template method
    void playGame() const {
        loadAssets();
        setup();
        startPlay();
        endPlay();
    }

protected:
    virtual void loadAssets() const = 0;
    virtual void setup() const = 0;
    virtual void startPlay() const = 0;

    void endPlay() const {
        std::cout << "Cleaning up resources...\n";
    }
};

// Concrete Class
class Chess : public Game {
protected:
    void loadAssets() const override {
        std::cout << "Loading chess assets...\n";
    }

    void setup() const override {
        std::cout << "Setting up chess board...\n";
    }

    void startPlay() const override {
        std::cout << "Starting chess game...\n";
    }
};

// Another Concrete Class
class Checkers : public Game {
protected:
    void loadAssets() const override {
        std::cout << "Loading checkers assets...\n";
    }

    void setup() const override {
        std::cout << "Setting up checkers board...\n";
    }

    void startPlay() const override {
        std::cout << "Starting checkers game...\n";
    }
};

int main() {
    Game* game = new Chess();
    game->playGame();
    delete game;

    game = new Checkers();
    game->playGame();
    delete game;

    return 0;
}


